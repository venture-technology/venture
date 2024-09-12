CREATE UNIQUE INDEX unique_child_currently_status
ON contracts (child_id)
WHERE status = 'currently';

