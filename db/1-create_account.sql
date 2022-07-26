-- Table: public.client_account

-- DROP TABLE IF EXISTS public.client_account;

CREATE TABLE IF NOT EXISTS public.client_account
(
    id bigint NOT NULL GENERATED ALWAYS AS IDENTITY ( INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 9223372036854775807 CACHE 1 ),
    email character varying(200) COLLATE pg_catalog."default" NOT NULL,
    mnemonic text COLLATE pg_catalog."default" NOT NULL,
    nickname character varying(30) COLLATE pg_catalog."default" NOT NULL,
    profile_image bytea NOT NULL,
    CONSTRAINT client_account_pkey PRIMARY KEY (id),
    CONSTRAINT unique_email UNIQUE (email),
    CONSTRAINT unique_mnemonic UNIQUE (mnemonic),
    CONSTRAINT unique_nickname UNIQUE (nickname)
)

    TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.client_account
    OWNER to postgres;