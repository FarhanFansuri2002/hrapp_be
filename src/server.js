import express from 'express'
import mysql from 'mysql2/promise'

const app = express()
const port = Number.parseInt(process.env.PORT || getServerPort(), 10)
const allowedOrigin = process.env.CLIENT_ORIGIN || 'http://localhost:5173'
const pool = mysql.createPool(getDatabaseConfig())

app.use((request, response, next) => {
  response.setHeader('Access-Control-Allow-Origin', allowedOrigin)
  response.setHeader('Access-Control-Allow-Headers', 'Content-Type, Authorization')
  response.setHeader('Access-Control-Allow-Methods', 'GET, OPTIONS')

  if (request.method === 'OPTIONS') {
    response.sendStatus(204)
    return
  }

  next()
})

app.get('/health', (_request, response) => {
  response.json({ status: 'ok' })
})

app.get('/api/v1/employees', async (_request, response) => {
  try {
    const [rows] = await pool.query(`
      SELECT id, name, role, department, status, joined_at AS joinedAt
      FROM employees
      ORDER BY name`)

    response.json(rows)
  } catch (error) {
    console.error('failed to load employees:', error.message)
    response.status(500).json({ error: 'failed to load employees' })
  }
})

app.listen(port, () => {
  console.log(`HR API listening on port ${port}`)
})

function getServerPort() {
  const address = process.env.SERVER_ADDRESS || ':8080'
  return address.startsWith(':') ? address.slice(1) : address.split(':').pop()
}

function getDatabaseConfig() {
  if (process.env.MYSQL_URL) {
    return process.env.MYSQL_URL
  }

  if (process.env.MYSQL_DSN) {
    return parseGoDsn(process.env.MYSQL_DSN)
  }

  return {
    host: process.env.MYSQL_HOST || '127.0.0.1',
    port: Number.parseInt(process.env.MYSQL_PORT || '3306', 10),
    user: process.env.MYSQL_USER || 'root',
    password: process.env.MYSQL_PASSWORD || '',
    database: process.env.MYSQL_DATABASE || 'hr_app'
  }
}

function parseGoDsn(dsn) {
  const match = dsn.match(/^([^:]+):([^@]*)@tcp\(([^:):]+)(?::(\d+))?\)\/([^?]+)/)
  if (!match) {
    throw new Error('MYSQL_DSN must use the Go MySQL format user:password@tcp(host:port)/database')
  }

  return {
    host: match[3],
    port: Number.parseInt(match[4] || '3306', 10),
    user: match[1],
    password: match[2],
    database: match[5]
  }
}
