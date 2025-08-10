CREATE TABLE todo_categories(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    create_at TIMESTAMP(3) WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP(3) WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,


    user_id TEXT NOT NULL,
    name TEXT NOT NULL,
    color TEXT DEFAULT '#6b7280',
    description TEXT

);


CREATE INDEX idx_todo_categories_user_id ON todo_categories(user_id);
CREATE UNIQUE INDEX todo_categories_unique_name ON todo_categories(user_id, name);


CREATE TRIGGER set_updated_at_todo_categories
    BEFORE UPDATE ON todo_categories
    FOR EACH ROW
    EXECUTE FUNCTION trigger_set_updated_at();




CREATE TABLE todos(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    create_at TIMESTAMP(3) WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP(3) WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,

    user_id TEXT NOT NULL,
    title TEXT NOT NULL,
    description TEXT,
    status TEXT NOT NULL DEFAULT 'draft',
    priority TEXT NOT NULL DEFAULT 'medium',
    due_date TIMESTAMP(3) WITH TIME ZONE,
    completed_at TIMESTAMP(3) WITH TIME ZONE,
    parent_todo_id UUID REFERENCES todos(id) ON DELETE SET NULL,
    category_id UUID REFERENCES todo_categories(id) ON DELETE SET NULL,
    metadata JSONB,
    sord_order SERIAL
);


CREATE INDEX idx_todos_user_id ON todos(user_id);
CREATE INDEX idx_todos_category_id ON todos(category_id);
CREATE INDEX idx_todos_parent_todo_id ON todos(parent_todo_id);
CREATE INDEX idx_todo_status ON todos(status);
CREATE INDEX idx_todo_priority ON todos(priority);
CREATE INDEX ids_todos_due_date ON todos(due_date);

CREATE TRIGGER set_updated_at_todos
    BEFORE UPDATE ON todos
    FOR EACH ROW
    EXECUTE FUNCTION trigger_set_updated_at();


CREATE TABLE todo_comments(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    create_at TIMESTAMP(3) WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP(3) WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,

    todo_id UUID NOT NULL REFERENCES todos ON DELETE CASCADE,
    user_id TEXT NOT NULL,
    content TEXT NOt NULL
);

CREATE INDEX idx_todos_comments_todo_id ON todo_comments(todo_id);
CREATE INDEX idx_todos_comments_user_id ON todo_comments(user_id);


CREATE TRIGGER set_updated_at_todo_comments
    BEFORE UPDATE ON todo_comments
    FOR EACH ROW
    EXECUTE FUNCTION trigger_set_updated_at();



-- Constraints for todos table
ALTER TABLE todos
ADD CONSTRAINT no_self_parent
CHECK (id != parent_todo_id);

CREATE INDEX idx_todos_hierarchy ON todos(parent_todo_id, sord_order);

CREATE INDEX idx_todos_user_status_priority ON todos(user_id, status, priority);
