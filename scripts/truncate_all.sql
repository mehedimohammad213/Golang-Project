-- Truncate all data and reset sequences (safe order with CASCADE)
-- Run: psql -h localhost -U postgres -d car_db -f scripts/truncate_all.sql
TRUNCATE
  installments,
  payment_history,
  order_items,
  orders,
  car_sub_details,
  car_grades,
  car_details,
  documents,
  car_photos,
  stocks,
  purchase_history,
  carts,
  cars,
  car_models,
  car_makes,
  permission_role,
  role_user,
  permissions,
  roles,
  users,
  rag_chunks
RESTART IDENTITY CASCADE;
