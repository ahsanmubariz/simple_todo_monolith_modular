DO $$
BEGIN
    IF EXISTS (SELECT 1 FROM information_schema.tables WHERE table_name = 'todos') THEN
        DROP TABLE todos;
    END IF;
END $$;