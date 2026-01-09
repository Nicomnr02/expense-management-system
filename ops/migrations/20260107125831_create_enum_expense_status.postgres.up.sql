DO $$
BEGIN
    IF NOT EXISTS (
        SELECT 1
        FROM pg_type
        WHERE typname = 'expense_status'
    ) THEN
        CREATE TYPE expense_status AS ENUM (
            'Awaiting Approval',
            'Pending',
            'Approved',
            'Rejected',
            'Auto-approved',
            'Under Review',
            'Completed'
        );
    END IF;
END$$;
