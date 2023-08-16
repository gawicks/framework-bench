import * as http from "http";
import postgres from "postgres";

const sql = postgres({ 
    username: 'admin',
    password: 'admin',
    database: 'db',
});
function getUsers() {
    const users = sql`SELECT * FROM USERS WHERE dob > DATE('1990-01-01')`;
    return users;
}
const server = http.createServer(async(req, res) => {
    const users = await getUsers();
    res.write(JSON.stringify(users));
    res.end();
})
server.listen(8000);