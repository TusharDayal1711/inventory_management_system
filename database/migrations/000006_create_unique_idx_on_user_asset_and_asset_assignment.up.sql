
--for user table
CREATE UNIQUE INDEX IF NOT EXISTS idx_users_contact_no_email
    ON users(contact_no, email);


--for asset table
CREATE UNIQUE INDEX IF NOT EXISTS idx_assets_serial_no
    ON assets(serial_no);


--for asset assignment
CREATE UNIQUE INDEX IF NOT EXISTS idx_asset_assignment_asset_id_employee_id
    ON asset_assign(asset_id, employee_id);

