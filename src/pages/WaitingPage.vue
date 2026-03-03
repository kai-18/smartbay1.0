<template>
  <q-page>
    <q-card>
      <q-card-section>
    <div class="row justify-between items-center">
      <div class="text-h6">Gestione Camion in attesa</div>
        <q-input
          v-model="searchQuery"
          label="Cerca..."
          dense
          outlined
          clearable
          debounce="300"
          class="q-ml-md"
          style="max-width: 300px;"
          >
        <template v-slot:prepend>
        <q-icon name="search" />
        </template>
        </q-input>
      </div>
    </q-card-section>

    <q-card-section>
      <q-btn label="Aggiungi" color="primary" @click="addDialog = true" />
    </q-card-section>

    <q-separator />

    <q-card-section>
      <q-table
        :rows="camion"
        :columns="columns"
        row-key="id"
        flat
        dense
        virtual-scroll
        :rows-per-page-options="[30, 50, 100]"
        class="scrollable-table"
        :filter="searchQuery"
      >
        <template v-slot:body-cell-actions="props">
        <q-td :props="props">
        <q-btn flat dense icon="edit" @click="openEditDialog(props.row)" />
        <q-btn flat dense icon="delete" color="negative" @click="showDeleteDialog(props.row)" />
        </q-td>
        </template>
      </q-table>
    </q-card-section>
    </q-card>
    <q-dialog v-model="addDialog">
      <q-card style="width: 500px;">
        <q-card-section>
          <div class="text-h6"></div>
        </q-card-section>

        <q-card-section>
          <q-form @submit.prevent="createWaiting">
            <q-select
              v-model="newWaiting.position"
              label="Tipologia"
              dense
              :options="positionOptions"
              emit-value
              map-options
              @update:model-value="() => handleCarico(newWaiting)"
            />
            <q-input
              v-if="newWaiting.position === 'Altro'"
              v-model="newWaiting.customPosition"
              label="Insierisci Tipologia"
              dense
              required
            />
            <q-input
              v-if="newWaiting.position === 'Carico'"
              v-model="newWaiting.customPosition"
              label="Insierisci Riferimenti Carico"
              dense
              required
            />
            <q-input v-model="newWaiting.name" label="Nome" dense />
            <q-input v-model="newWaiting.lastname" label="Cognome" dense />
            <q-input v-model="newWaiting.username" label="Telefono" dense />
            <q-input v-model="newWaiting.email" label="Targa" dense />

            <q-card-actions align="right">
              <q-btn flat label="Annulla" color="negative" @click="addDialog = false" />
              <q-btn type="submit" label="Salva" color="primary" />
            </q-card-actions>
          </q-form>
        </q-card-section>
      </q-card>
    </q-dialog>
    <q-dialog v-model="editDialog">
      <q-card>
        <q-card-section>
          <div class="text-h6">Modifica</div>
        </q-card-section>

        <q-card-section class="editCamion">
          <q-form @submit.prevent="updateCamion">
            <q-input v-model="editCamion.name" label="Nome" dense />
            <q-input v-model="editCamion.lastname" label="Cognome" dense />
            <q-input v-model="editCamion.username" label="Telefono" dense />
            <q-input v-model="editCamion.email" label="Targa" dense />
            <q-select
              v-model="editCamion.position"
              label="Tipologia"
              dense
              :options="positionOptions"
              emit-value
              map-options
              @update:model-value="() => handleCambioTipologia(editCamion)"
            />

            <q-input
              v-if="editCamion.position === 'Altro'"
              v-model="editCamion.customPosition"
              label="Insierisci Tipologia"
              dense
            />
            <q-card-actions align="right">
              <q-btn flat label="Cancel" color="negative" @click="editDialog = false" />
              <q-btn type="submit" label="Update" color="primary" />
            </q-card-actions>
          </q-form>
        </q-card-section>
      </q-card>
    </q-dialog>
    <q-dialog v-model="deleteDialog">
      <q-card>
        <q-card-section>
          <div class="text-h6">Sei sicuro di voler eliminare?</div>
        </q-card-section>

        <q-card-actions align="right">
          <q-btn flat label="Cancel" color="negative" @click="deleteDialog = false" />
          <q-btn label="Delete" color="primary" @click="confirmDelete" />
        </q-card-actions>
      </q-card>
    </q-dialog>
  </q-page>
</template>


<script setup>
import { ref, onMounted } from 'vue'
import axios from 'axios'

const camion = ref([])
const addDialog = ref(false)
const editDialog = ref(false)
const searchQuery = ref('')
const deleteDialog = ref(false)
const camionToDelete = ref(null)
const positionOptions = [
  { label: 'Carico', value: 'Carico' },
  { label: 'Scarico', value: 'Scarico' },
  { label: 'Navettaggio', value: 'Navettaggio' },
  { label: 'Trasferimento', value: 'Trasferimento' },
  { label: 'Reso', value: 'Reso' },
  { label: 'Altro', value: 'Altro' }
]

function handleCambioTipologia(camion) {
  if (camion.position !== 'Altro') {
    camion.customPosition = ''
  }
}
function handleCarico(camion) {
  if (camion.tipo !== 'Carico') {
    camion.customTipo = ''
  }
}
function openEditDialog(camion) {
  editCamion.value = { ...camion }
  editDialog.value = true
}

function showDeleteDialog(camion) {
  camionToDelete.value = camion
  deleteDialog.value = true
}

function confirmDelete() {
  deleteCamion(camionToDelete.value.id)
  deleteDialog.value = false
}
const newWaiting = ref({
  tipologia: '',
  nome: '',
  cognome: '',
  telefono: '',
  targa: '',
})

const editCamion = ref({})

const columns = [
  // {
  //   name: 'id',
  //   label: 'ID',
  //   field: row => row.id,
  //   sortable: true,
  //   align: 'center'
  // },
  {
    name: 'tipologia',
    label: 'Tipologia',
    field: row => row.tipologia,
    sortable: true,
    align: 'left'
  },
  {
    name: 'nome',
    label: 'Nome',
    field: row => row.nome,
    sortable: true,
    align: 'left'
  },
  {
    name: 'cognome',
    label: 'Cognome',
    field: row => row.cognome,
    sortable: true,
    align: 'left'
  },
  {
    name: 'telefono',
    label: 'Telefono',
    field: row => row.telefono,
    sortable: true,
    align: 'left'
  },
  {
    name: 'targa',
    label: 'Targa',
    field: row => row.targa,
    sortable: true,
    align: 'left'
  },
  {
    name: 'azioni',
    label: 'Azioni',
    field: 'azioni'
  }
]

async function fetchCamion() {
  try {
    const response = await axios.get('IP ADDRESS HERE')
    camion.value = response.data
  } catch (error) {
    console.error('Error fetching camion:', error)
  }
}

async function createWaiting() {
  try {
    const payload = {
      ...newWaiting.value,
      position: newWaiting.value.position === 'Other' ? newWaiting.value.customPosition : newWaiting.value.position
    }

    await axios.post('http://192.168.1.101:8080/api/camion/create', payload)

    newWaiting.value = { nome: '', Cognome: '', Telefono: '', Targa: ''}
    addDialog.value = false
    fetchCamion()
  } catch (error) {
    console.error('Error creating camion:', error)
  }
}

async function updateCamion() {
  try {
    const payload = {
      ...editCamion.value,
      tipologia: editCamion.value.tipologia === 'Altro' ? editCamion.value.customTipologia : editCamion.value.tipologia
    }

    await axios.put('http://192.168.1.101:8080/api/camion/update', payload)

    editDialog.value = false
    fetchCamion()
  } catch (error) {
    console.error('Error updating camion:', error)
  }
}

async function deleteCamion(id) {
  try {
    await axios.delete('http://192.168.1.101:8080/api/camion/delete', { data: { id: id } })
    fetchCamion()
  } catch (error) {
    console.error('Error deleting camion:', error)
  }
}



onMounted(() => {
  fetchCamion()
})
</script>

<style scoped>
.scrollable-table {
  max-height: 600px;  /* Set height for scrollable container */
  overflow-y: auto;   /* Allow vertical scrolling */
}

.q-table__header {
  position: sticky;  /* Sticky header */
  top: 0;
  background-color: white;
  z-index: 1;  /* Ensures header is always on top of the body */
}

.q-th {
  background-color: #f5f5f5;  /* Optional: header background color */
}

.q-mt-sm {
  margin-top: 8px;
}
.editCamion{
  width: 400px;
}
</style>
