import cluster from "cluster"
import * as http from "http"
import postgres from "postgres";
const numCPUs = 4;

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
  for (let i = 0; i < numCPUs; i++) {
    cluster.fork();
  }

  // Listen for worker exit event and replace it
  cluster.on('exit', (worker, code, signal) => {
    console.log(`Worker ${worker.process.pid} died`);
    cluster.fork(); // Create a new worker to replace the one that died
  });
} else {
  // Each worker handles its own HTTP server
  http.createServer(async (req, res) => {
    const users = await getUsers();
    res.write(JSON.stringify(users));
    res.end();
  }).listen(8001);

  console.log(`Worker ${cluster.worker.id} started`);
}
