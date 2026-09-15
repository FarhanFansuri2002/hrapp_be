CREATE TABLE IF NOT EXISTS employees (
    id VARCHAR(36) PRIMARY KEY,
    name VARCHAR(120) NOT NULL,
    role VARCHAR(120) NOT NULL,
    department VARCHAR(120) NOT NULL,
    status ENUM('Aktif', 'Cuti', 'Nonaktif') NOT NULL DEFAULT 'Aktif',
    joined_at DATE NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

INSERT INTO employees (id, name, role, department, status, joined_at)
VALUES
    ('EMP-001', 'Aulia Sari', 'People Operations', 'People', 'Aktif', '2024-03-11'),
    ('EMP-002', 'Raka Pratama', 'Frontend Engineer', 'Engineering', 'Aktif', '2023-08-21'),
    ('EMP-003', 'Alya Rahma', 'Product Manager', 'Product', 'Cuti', '2022-11-07')
ON DUPLICATE KEY UPDATE name = VALUES(name);