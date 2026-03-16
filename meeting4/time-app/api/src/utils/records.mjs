import pool from './pgPool.mjs'

const readRecords = async () => {
  const result = await pool.query('SELECT * FROM times ORDER BY created_at DESC')
  return result.rows
}

const insertRecord = async (time) => {
  const result = await pool.query('INSERT INTO times (time) VALUES ($1) RETURNING *', [time])
  console.log(`New time ${time} was saved to the DB`)
  return result.rows[0]
}

const deleteRecord = async (id) => {
  const result = await pool.query('DELETE FROM times WHERE id=$1', [id])
  console.log(`Time with id ${id} was deleted from the DB`)
  return result
}

export { readRecords, insertRecord, deleteRecord }
