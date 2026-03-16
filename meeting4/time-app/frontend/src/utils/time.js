import moment from 'moment'

const API_URL = import.meta.env.VITE_API_URL || 'http://localhost:5000'

function startInterval() {
  setInterval(() => {
    this.currentTime = moment().format('HH:mm:ss')
  }, 1000)
}

async function saveTime() {
  const time = this.currentTime
  const res = await fetch(`${API_URL}/times`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify({ time }),
  })
  const json = await res.json()
  if (json.id) {
    this.savedTimes.unshift({ id: json.id, time })
    this.$toast.success(`Time ${time} saved`, { position: 'top-right' })
  }
}

async function deleteTime(id) {
  const res = await fetch(`${API_URL}/time/${id}`, {
    method: 'DELETE',
  })
  const json = await res.json()
  if (json.rowCount) {
    this.savedTimes = this.savedTimes.filter((savedTime) => savedTime.id !== id)
    this.$toast.error(`Time with ID ${id} was removed`, {
      position: 'top-right',
    })
  }
}

export { startInterval, saveTime, deleteTime }
