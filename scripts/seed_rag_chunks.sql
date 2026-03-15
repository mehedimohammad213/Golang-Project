-- Seed: 5 rows for rag_chunks (pgvector must be enabled)
-- Run: psql -h localhost -U postgres -d car_db -f scripts/seed_rag_chunks.sql

-- Build a 1536-dimensional zero vector (required by embedding column)
WITH zeros AS (
  SELECT ('[' || string_agg('0.0', ',') || ']')::vector AS emb
  FROM generate_series(1, 1536)
),
data AS (
  SELECT * FROM (VALUES
    ('car', '1', 'Toyota Camry 2024. Sedan, Petrol, Automatic. Black, 15000 km. Premium package.'),
    ('car', '2', 'Honda Accord 2023. Sedan, Hybrid, CVT. White, 8000 km. Full options.'),
    ('car', '3', 'Nissan Altima 2024. Sedan, Petrol, Automatic. Silver, 12000 km.'),
    ('car', '4', 'Mazda 6 2023. Sedan, Petrol, Automatic. Red, 5000 km. Sport package.'),
    ('car', '5', 'Subaru Outback 2024. Wagon, AWD, Automatic. Blue, 10000 km. Family SUV.')
  ) AS t(source_type, source_id, content)
)
INSERT INTO rag_chunks (source_type, source_id, content, embedding, metadata)
SELECT d.source_type, d.source_id, d.content, z.emb, '{}'::jsonb
FROM data d
CROSS JOIN zeros z;
