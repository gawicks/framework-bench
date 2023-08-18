import * as http from "http";
import postgres from "postgres";
import prexit from 'prexit';

const sql = postgres({
    username: 'admin',
    password: 'admin',
    database: 'db',
});
function getUsers() {
    const users = sql`SELECT * FROM USERS WHERE dob > '1990-01-01'`;
    return users;
}
const server = http.createServer(async (req, res) => {
    const users = await getUsers();
    res.end(JSON.stringify(users));
})
server.listen(8000);
console.log("Server started");
prexit(async () => {
    await sql.end({ timeout: 5 });
    await server.close();
    console.log("EXIT CALLED");
})
