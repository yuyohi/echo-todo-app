CREATE TABLE IF NOT EXISTS task (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    title TEXT NOT NULL,
    description TEXT,
    status TEXT DEFAULT 'todo',
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    CHECK (status IN ('todo', 'completed'))
);
