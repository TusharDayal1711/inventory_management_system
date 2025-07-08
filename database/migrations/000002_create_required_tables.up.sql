--user table
CREATE TABLE IF NOT EXISTS users(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    username TEXT NOT NULL,
    email TEXT NOT NULL,
    contact_no TEXT,
    system_role system_role DEFAULT NULL,
    employee_type employee_type DEFAULT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT now(),
    archived_at TIMESTAMP
);

--assets table
CREATE TABLE IF NOT EXISTS assets(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    brand TEXT NOT NULL,
    model TEXT NOT NULL,
    serial_no TEXT NOT NULL,
    purchase_date timestamp,
    owned_by ownership NOT NULL DEFAULT 'remotestate',
    type asset_type,
    warranty_start TIMESTAMP,
    warranty_expire TIMESTAMP,
    added_by UUID REFERENCES users(id),
    added_at TIMESTAMP WITH TIME ZONE DEFAULT now(),
    status asset_status NOT NULL DEFAULT 'available',
    archived_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS asset_assign(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    asset_id UUID NOT NULL REFERENCES assets(id),
    employee_id UUID NOT NULL REFERENCES users(id),
    assigned_at TIMESTAMP DEFAULT now(),
    returned_at TIMESTAMP,
    return_reason TEXT,
    archived_at TIMESTAMP
);



