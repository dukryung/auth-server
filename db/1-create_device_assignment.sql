-- Table: public.device_assignment

-- DROP TABLE IF EXISTS public.device_assignment;

CREATE TABLE IF NOT EXISTS public.device_assignment
(
    mnemonic text COLLATE pg_catalog."default" NOT NULL,
    email character varying(200) COLLATE pg_catalog."default" NOT NULL,
    device_id character varying(100) COLLATE pg_catalog."default" NOT NULL,
    CONSTRAINT device_id UNIQUE (device_id)
        INCLUDE(device_id),
    CONSTRAINT email UNIQUE (email),
    CONSTRAINT mnemonic UNIQUE (mnemonic)
)

    TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.device_assignment
    OWNER to postgres;