create sequence seq_user;

create table tb_user (
	user_id int not null default nextval('seq_user'), 
	user_role int,
	user_name varchar(50) not null,
	user_email varchar(100) not null,
	user_naverid varchar(100),
	user_kakaoid varchar(100),
	user_updateat timestamp default now()
)

INSERT INTO public.tb_user (user_role, user_name, user_email, user_naverid) VALUES (0, '%s', '%s', '%s')

SELECT user_role, user_name, user_email, user_naverid FROM public.tb_user WHERE user_role > 9

SELECT user_id FROM public.tb_user WHERE user_email='%s' and user_naverid='%s'