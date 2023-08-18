using System.Text.Json;
using Npgsql;

var builder = WebApplication.CreateBuilder(args);
var app = builder.Build();

var connString = "Host=localhost;Username=admin;Password=admin;Database=db;Maximum Pool Size=12;NoResetOnClose=true;Enlist=false;Max Auto Prepare=4;Multiplexing=true;Write Coalescing Buffer Threshold Bytes=1000";

app.MapGet("/", async () =>
{
    await using var conn = new NpgsqlConnection(connString);
    await conn.OpenAsync();

    var users = new List<User>();
    await using (var cmd = new NpgsqlCommand("SELECT id, name, dob, username, gender, email, phone FROM USERS WHERE dob > '1990-01-01'", conn))
    await using (var reader = await cmd.ExecuteReaderAsync())
    {
        while (await reader.ReadAsync())
        {
            var user = new User
            {
                Id = reader.GetInt32(0),
                Name = reader.GetString(1),
                Dob = reader.GetDateTime(2),
                Username = reader.GetString(3),
                Gender = reader.GetString(4),
                Email = reader.GetString(5),
                Phone = reader.GetString(6)
            };
            users.Add(user);     
        }
    }
    var json = JsonSerializer.Serialize(users);
    return json;
});

app.Run("http://localhost:8004");

public class User
{
    public int Id { get; set; }
    public string Name { get; set; }
    public DateTime Dob { get; set; }
    public string Username { get; set; }
    public string Phone { get; set; }
    public string Email { get; set; }
    public string Gender { get; set; }


}

