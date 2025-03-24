create table scrap_papers (
	id uuid primary key default gen_random_uuid(),
	content text not null,
	is_private boolean not null default false,
	created_at timestamp with time zone default now(),
	updated_at timestamp with time zone default now()
);