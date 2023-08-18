import cluster from "cluster"
import * as http from "http"
import postgres from "postgres";
import prexit from 'prexit'

const workers = 4;

const sql = postgres({ 
    username: 'admin',
    password: 'admin',
    database: 'db',
    max_lifetime: 1
});
function getUsers() {
    const users = sql`SELECT * FROM USERS WHERE dob > '1990-01-01'`;
    return users;
}
if (cluster.isPrimary) {
  for (let i = 0; i < workers; i++) {
    cluster.fork();
  }
} else {
  const server = http.createServer(async (req, res) => {
    const users = await getUsers();
    res.end(JSON.stringify(users));
  }).listen(8001);
  prexit(async () => {
    await sql.end({ timeout: 5 })
    await server.close()
    console.log("EXIT CALLED")
  })

  console.log(`Worker ${cluster.worker.id} started`);  
}

