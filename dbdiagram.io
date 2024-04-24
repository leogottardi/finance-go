// Use DBML to define your database structure
// Docs: https://dbml.dbdiagram.io/docs

Table customers {
  id string [primary key]
  name string
  document string
  created_at timestamp [default: `now()`]
  updated_at timestamp 
}

Table accounts {
  id string [primary key]
  customer_id varchar
  due_day int
  limit int
  created_at timestamp [default: `now()`]
  updated_at timestamp
}

Table transactions {
  id string [primary key]
  amount int
  type string
  created_at timestamp [default: `now()`]
  updated_at timestamp
}

Table invoices {
  id string [primary key]
  account_id string
  status string
  amount int
  term int
  created_at timestamp [default: `now()`]
  updated_at timestamp
}

Table transaction_invoice {
  id string [primary key]
  transaction_id string
  invoice_id string
  created_at timestamp [default: `now()`]
  updated_at timetsmp
}

Ref: accounts.customer_id - customers.id 

Ref: transaction_invoice.invoice_id > invoices.id
Ref: transaction_invoice.transaction_id > transactions.id

Ref: invoices.account_id > accounts.id
