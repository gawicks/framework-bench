import cluster from "cluster"
import * as http from "http"
import postgres from "postgres";
const workers = 4;

const sql = postgres({ 
    username: 'admin',
    password: 'admin',
    database: 'db',
});
function getUsers() {
    const users = sql`SELECT * FROM USERS WHERE dob > DATE('1990-01-01')`;
    return users;
}
if (cluster.isPrimary) {
  // Fork workers for each CPU core
  for (let i = 0; i < workers; i++) {
    cluster.fork();
  }
} else {
  http.createServer(async (req, res) => {
    const users = await getUsers();
    res.write(JSON.stringify(users));
    res.end();
  }).listen(8001);

  console.log(`Worker ${cluster.worker.id} started`);
}
