import pkg from 'pg'
const { Pool } = pkg

// const DB_HOST = process.env.DB_HOST || 'db'
// const DB_USER = process.env.DB_USER || 'root'
// const DB_PORT = process.env.DB_PORT || '5432'
// const DB_PASSWORD = process.env.DB_PASSWORD || 'password'
// const DB_NAME = process.env.DB_NAME || 'time_db'

// console.log(process.env)

const pool = new Pool({
  host: 'db', // instead of postgresql we use service name from docker compose configuration file
  port: '5432',
  user: 'root',
  password: 'password',
  database: 'time_db',
})

const CREATE_TIMES_TABLE_SQL = `CREATE TABLE IF NOT EXISTS times (
  id SERIAL PRIMARY KEY,
  time TEXT,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
)`

pool.connect((err, client, release) => {
  if (!err) {
    console.log('Connected to the PostgreSQL DB')
    client.query(CREATE_TIMES_TABLE_SQL, (err) => {
      if (!err) {
        console.log('Times table was created')
      }
      release()
    })
  }
})

export default pool
