SELECT a.id, a.username, b.username as parentusername from public.user a left join public.user b on a.parent = b.id
