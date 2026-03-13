# Exercises

Load provided test data.

```
psql -h 127.0.0.1 -p 5432 -U user postgres < data.sql
```

You should now have new schema `exercise1` with data for this exercise.

You should have a table with columns:

- `id` (integer, primary key)
- `name` (text, not null)
- `age` (integer)
- `grade` (text)

# Tasks
Remark: Pay attention, that in seeding script we are using schema. So, if you want to query to specific table from the schema, you should the following formula schema_name.table_name (ex. FROM excercise1.students)

List all students.
List all students older than 20.
List all students with grade 'B'.
List all students whose name starts with 'A'.
List all students without a grade.

Update 'Bob' age to 21.
Update 'Damian' grade to 'C'.

Add new student `Gandalf` with age 23 and grade 'A'

Create new table `courses` with columns:
- `id` (integer, primary key)
- `student_id` (integer, references students table)
- `course_name` (text, not null)
- `credits` (integer)

Add following course records:
- Alice (student_id=1) is taking "Mathematics" (3 credits)
- Alice (student_id=1) is taking "Physics" (4 credits)
- Bob (student_id=2) is taking "Chemistry" (3 credits)

List the sum of all credits for each student

Delete all students who have grade 'D' from students table.

# Troubleshooting

If you have problem to connect to "user", you can try to create a new one:
1. sudo -u postgres psql 
2. CREATE USER demo WITH PASSWORD 'Secret1!';
3. CREATE DATABASE demo_db;
4. GRANT ALL PRIVILEGES ON DATABASE demo_db TO demo;

Seed db:
psql -h 127.0.0.1 -p 5432 -U demo demo_db < data.sql

And login to new database with new role:
psql -U demo -p 5433 -h localhost -d demo_db  