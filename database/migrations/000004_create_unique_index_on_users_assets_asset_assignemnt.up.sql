--user
CREATE UNIQUE INDEX IF NOT EXISTS idx_users_email
    ON users(email)
    WHERE archived_at IS NULL;

CREATE UNIQUE INDEX IF NOT EXISTS idx_users_contact_no
    ON users(contact_no)
    WHERE archived_at IS NULL;

--assets
CREATE UNIQUE INDEX IF NOT EXISTS idx_assets_serial_no_active
    ON assets(serial_no)
    WHERE archived_at IS NULL;

--asset_assign
CREATE UNIQUE INDEX IF NOT EXISTS idx_asset_assign_assign_asset_id_employee_id
    ON asset_assign(asset_id, employee_id)
    WHERE archived_at IS NULL AND returned_at IS NULL;

