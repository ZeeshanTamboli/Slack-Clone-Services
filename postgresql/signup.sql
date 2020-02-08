/* Creating users table */
create table users(id serial primary key, first_name varchar(100), last_name varchar(100), email varchar(100) unique, created_at timestamptz default now(), updated_at timestamptz default now());


/* Creating Workspaces table */
create table workspaces(id serial primary key, name varchar(100) unique, owner_id int, foreign key(owner_id) references users(id), created_at timestamptz default now(), updated_at timestamptz default now());


/* Update timestamp when row is updated in PostgreSQL */
CREATE OR REPLACE FUNCTION update_update_at_column()
RETURNS TRIGGER AS $$
BEGIN
   NEW.updated_at = now(); 
   RETURN NEW;
END;
$$ language 'plpgsql';

CREATE TRIGGER update_users_updated_at BEFORE UPDATE
ON users FOR EACH ROW EXECUTE PROCEDURE 
update_update_at_column();

CREATE TRIGGER update_workspaces_updated_at BEFORE UPDATE
ON workspaces FOR EACH ROW EXECUTE PROCEDURE 
update_update_at_column();



/* Creating our bridge table */
create table users_workspaces (user_id int, workspace_id int, foreign key(user_id) references users(id), foreign key(workspace_id) references workspaces(id), primary key(user_id, workspace_id));