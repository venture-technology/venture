--
-- PostgreSQL database dump
--

-- Dumped from database version 12.20 (Ubuntu 12.20-0ubuntu0.20.04.1)
-- Dumped by pg_dump version 16.8 (Ubuntu 16.8-0ubuntu0.24.04.1)

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: public; Type: SCHEMA; Schema: -; Owner: postgres
--

-- *not* creating schema, since initdb creates it


ALTER SCHEMA public OWNER TO postgres;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: children; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.children (
    id integer NOT NULL,
    name character varying(100) NOT NULL,
    rg character varying(9) NOT NULL,
    responsible_id character varying(11) NOT NULL,
    shift text NOT NULL
);


ALTER TABLE public.children OWNER TO postgres;

--
-- Name: children_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.children_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.children_id_seq OWNER TO postgres;

--
-- Name: children_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.children_id_seq OWNED BY public.children.id;


--
-- Name: contracts; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.contracts (
    id integer NOT NULL,
    uuid text NOT NULL,
    status text NOT NULL,
    stripe_subscription_id text NOT NULL,
    stripe_price_id text NOT NULL,
    stripe_product_id text NOT NULL,
    signing_url text NOT NULL,
    driver_cnh text NOT NULL,
    school_cnpj text NOT NULL,
    kid_rg text NOT NULL,
    responsible_cpf text NOT NULL,
    created_at bigint DEFAULT date_part('epoch'::text, now()) NOT NULL,
    updated_at bigint DEFAULT date_part('epoch'::text, now()) NOT NULL,
    expire_at bigint NOT NULL,
    amount numeric(10,2) NOT NULL,
    anual_amount numeric(10,2) NOT NULL,
    CONSTRAINT contracts_status_check CHECK ((status = ANY (ARRAY['currently'::text, 'canceled'::text, 'expired'::text])))
);


ALTER TABLE public.contracts OWNER TO postgres;

--
-- Name: contracts_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.contracts_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.contracts_id_seq OWNER TO postgres;

--
-- Name: contracts_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.contracts_id_seq OWNED BY public.contracts.id;


--
-- Name: drivers; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.drivers (
    id integer NOT NULL,
    name character varying(100) NOT NULL,
    cpf character varying(14) NOT NULL,
    email character varying(100) NOT NULL,
    password character varying(100) NOT NULL,
    cnh character varying(20) NOT NULL,
    qr_code character varying(100) NOT NULL,
    amount numeric(10,2) NOT NULL,
    street character varying(100) NOT NULL,
    number character varying(10) NOT NULL,
    complement character varying(10),
    zip character varying(8) NOT NULL,
    phone text NOT NULL,
    pix_key character varying(100),
    municipal_record text NOT NULL,
    profile_image text DEFAULT 'null'::text,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    car_name character varying(225),
    car_year character varying(4),
    car_capacity character varying(3),
    schedule character varying(3),
    seats_remaining integer,
    seats_morning integer,
    seats_afternoon integer,
    seats_night integer,
    state character varying(100) DEFAULT ''::character varying NOT NULL,
    city character varying(100) DEFAULT ''::character varying NOT NULL,
    accessibility boolean DEFAULT false,
    biography character varying(1000) DEFAULT ''::character varying,
    descriptions character varying(550) DEFAULT ''::character varying,
    neighborhood character varying(255) DEFAULT ''::character varying NOT NULL
    seats_version bigint DEFAULT 0 NOT NULL
);


ALTER TABLE public.drivers OWNER TO postgres;

--
-- Name: drivers_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.drivers_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.drivers_id_seq OWNER TO postgres;

--
-- Name: drivers_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.drivers_id_seq OWNED BY public.drivers.id;


--
-- Name: invites; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.invites (
    requester character varying(14),
    guester character varying(20),
    status text NOT NULL,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    accepted_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    id integer NOT NULL
);


ALTER TABLE public.invites OWNER TO postgres;

--
-- Name: invites_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.invites_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.invites_id_seq OWNER TO postgres;

--
-- Name: invites_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.invites_id_seq OWNED BY public.invites.id;


--
-- Name: kids; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.kids (
    id integer NOT NULL,
    name character varying(100) NOT NULL,
    rg character varying(9) NOT NULL,
    responsible_id character varying(11) NOT NULL,
    shift text NOT NULL,
    profile_image text,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    attendance_permission boolean DEFAULT true
);


ALTER TABLE public.kids OWNER TO postgres;

--
-- Name: kids_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.kids_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.kids_id_seq OWNER TO postgres;

--
-- Name: kids_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.kids_id_seq OWNED BY public.kids.id;


--
-- Name: partners; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.partners (
    id integer NOT NULL,
    driver_cnh character varying(20) NOT NULL,
    school_cnpj character varying(14) NOT NULL,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP
);


ALTER TABLE public.partners OWNER TO postgres;

--
-- Name: partners_record_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.partners_record_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.partners_record_seq OWNER TO postgres;

--
-- Name: partners_record_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.partners_record_seq OWNED BY public.partners.id;


--
-- Name: payouts; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.payouts (
    id uuid NOT NULL,
    payout json NOT NULL,
    route text NOT NULL
);


ALTER TABLE public.payouts OWNER TO postgres;

--
-- Name: responsible; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.responsible (
    id integer NOT NULL,
    name character varying(100) NOT NULL,
    cpf character varying(11) NOT NULL,
    email character varying(100) NOT NULL,
    password character varying(100) NOT NULL,
    street character varying(100) NOT NULL,
    number text NOT NULL,
    complement text,
    zip character varying(8) NOT NULL,
    status text NOT NULL,
    card_token text,
    payment_method_id text,
    customer_id text NOT NULL,
    phone text NOT NULL
);


ALTER TABLE public.responsible OWNER TO postgres;

--
-- Name: responsibles; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.responsibles (
    id integer NOT NULL,
    name character varying(100) NOT NULL,
    cpf character varying(11) NOT NULL,
    email character varying(100) NOT NULL,
    password character varying(100) NOT NULL,
    street character varying(100) NOT NULL,
    number text NOT NULL,
    complement text,
    zip character varying(8) NOT NULL,
    card_token text NOT NULL,
    payment_method_id text,
    customer_id text NOT NULL,
    phone text NOT NULL,
    profile_image text DEFAULT 'null'::text,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    state character varying(100) DEFAULT ''::character varying NOT NULL,
    city character varying(100) DEFAULT ''::character varying NOT NULL,
    neighborhood character varying(255) DEFAULT ''::character varying NOT NULL
);


ALTER TABLE public.responsibles OWNER TO postgres;

--
-- Name: responsible_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.responsible_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.responsible_id_seq OWNER TO postgres;

--
-- Name: responsible_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.responsible_id_seq OWNED BY public.responsibles.id;


--
-- Name: responsible_id_seq1; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.responsible_id_seq1
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.responsible_id_seq1 OWNER TO postgres;

--
-- Name: responsible_id_seq1; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.responsible_id_seq1 OWNED BY public.responsible.id;


--
-- Name: schema_migrations; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.schema_migrations (
    version bigint NOT NULL,
    dirty boolean NOT NULL
);


ALTER TABLE public.schema_migrations OWNER TO postgres;

--
-- Name: schools; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.schools (
    id integer NOT NULL,
    name character varying(100) NOT NULL,
    password character varying(100) NOT NULL,
    cnpj character varying(14) NOT NULL,
    street character varying(100) NOT NULL,
    number character varying(10) NOT NULL,
    zip character varying(8) NOT NULL,
    email character varying(100) NOT NULL,
    complement character varying(10),
    phone text NOT NULL,
    profile_image text DEFAULT 'null'::text,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    state character varying(100) DEFAULT ''::character varying NOT NULL,
    city character varying(100) DEFAULT ''::character varying NOT NULL,
    neighborhood character varying(255) DEFAULT ''::character varying NOT NULL
);


ALTER TABLE public.schools OWNER TO postgres;

--
-- Name: schools_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.schools_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.schools_id_seq OWNER TO postgres;

--
-- Name: schools_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.schools_id_seq OWNED BY public.schools.id;


--
-- Name: temp_contracts; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.temp_contracts (
    id integer NOT NULL,
    uuid character varying(36) NOT NULL,
    signing_url text NOT NULL,
    created_at bigint NOT NULL,
    expired_at bigint NOT NULL,
    status character varying(50) NOT NULL,
    driver_cnh character varying(20) NOT NULL,
    responsible_cpf character varying(14) NOT NULL,
    school_cnpj character varying(18) NOT NULL,
    kid_rg character varying(12) NOT NULL,
    driver_signed_at bigint,
    responsible_signed_at bigint,
    signature_request_id text NOT NULL
);


ALTER TABLE public.temp_contracts OWNER TO postgres;

--
-- Name: temp_contracts_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.temp_contracts_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.temp_contracts_id_seq OWNER TO postgres;

--
-- Name: temp_contracts_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.temp_contracts_id_seq OWNED BY public.temp_contracts.id;


--
-- Name: children id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.children ALTER COLUMN id SET DEFAULT nextval('public.children_id_seq'::regclass);


--
-- Name: contracts id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.contracts ALTER COLUMN id SET DEFAULT nextval('public.contracts_id_seq'::regclass);


--
-- Name: drivers id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.drivers ALTER COLUMN id SET DEFAULT nextval('public.drivers_id_seq'::regclass);


--
-- Name: invites id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.invites ALTER COLUMN id SET DEFAULT nextval('public.invites_id_seq'::regclass);


--
-- Name: kids id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.kids ALTER COLUMN id SET DEFAULT nextval('public.kids_id_seq'::regclass);


--
-- Name: partners id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.partners ALTER COLUMN id SET DEFAULT nextval('public.partners_record_seq'::regclass);


--
-- Name: responsible id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.responsible ALTER COLUMN id SET DEFAULT nextval('public.responsible_id_seq1'::regclass);


--
-- Name: responsibles id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.responsibles ALTER COLUMN id SET DEFAULT nextval('public.responsible_id_seq'::regclass);


--
-- Name: schools id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.schools ALTER COLUMN id SET DEFAULT nextval('public.schools_id_seq'::regclass);


--
-- Name: temp_contracts id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.temp_contracts ALTER COLUMN id SET DEFAULT nextval('public.temp_contracts_id_seq'::regclass);


--
-- Name: children children_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.children
    ADD CONSTRAINT children_pkey PRIMARY KEY (rg);


--
-- Name: contracts contracts_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.contracts
    ADD CONSTRAINT contracts_pkey PRIMARY KEY (id);


--
-- Name: contracts contracts_uuid_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.contracts
    ADD CONSTRAINT contracts_uuid_key UNIQUE (uuid);


--
-- Name: drivers drivers_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.drivers
    ADD CONSTRAINT drivers_pkey PRIMARY KEY (cnh);


--
-- Name: invites invites_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.invites
    ADD CONSTRAINT invites_pkey PRIMARY KEY (id);


--
-- Name: partners partners_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.partners
    ADD CONSTRAINT partners_pkey PRIMARY KEY (id);


--
-- Name: payouts payouts_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.payouts
    ADD CONSTRAINT payouts_pkey PRIMARY KEY (id);


--
-- Name: responsibles responsible_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.responsibles
    ADD CONSTRAINT responsible_pkey PRIMARY KEY (cpf);


--
-- Name: responsible responsible_pkey1; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.responsible
    ADD CONSTRAINT responsible_pkey1 PRIMARY KEY (cpf);


--
-- Name: schema_migrations schema_migrations_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.schema_migrations
    ADD CONSTRAINT schema_migrations_pkey PRIMARY KEY (version);


--
-- Name: schools schools_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.schools
    ADD CONSTRAINT schools_pkey PRIMARY KEY (cnpj);


--
-- Name: temp_contracts temp_contracts_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.temp_contracts
    ADD CONSTRAINT temp_contracts_pkey PRIMARY KEY (id);


--
-- Name: temp_contracts temp_contracts_uuid_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.temp_contracts
    ADD CONSTRAINT temp_contracts_uuid_key UNIQUE (uuid);


--
-- Name: drivers unique_cnh; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.drivers
    ADD CONSTRAINT unique_cnh UNIQUE (cnh);


--
-- Name: schools unique_cnpj; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.schools
    ADD CONSTRAINT unique_cnpj UNIQUE (cnpj);


--
-- Name: responsibles unique_cpf; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.responsibles
    ADD CONSTRAINT unique_cpf UNIQUE (cpf);


--
-- Name: invites unique_invite; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.invites
    ADD CONSTRAINT unique_invite UNIQUE (requester, guester);


--
-- Name: kids unique_kid; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.kids
    ADD CONSTRAINT unique_kid PRIMARY KEY (rg);


--
-- Name: partners unique_partner; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.partners
    ADD CONSTRAINT unique_partner UNIQUE (driver_cnh, school_cnpj);


--
-- Name: children children_responsible_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.children
    ADD CONSTRAINT children_responsible_id_fkey FOREIGN KEY (responsible_id) REFERENCES public.responsible(cpf) ON DELETE CASCADE;


--
-- Name: contracts fk_driver; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.contracts
    ADD CONSTRAINT fk_driver FOREIGN KEY (driver_cnh) REFERENCES public.drivers(cnh) ON DELETE CASCADE;


--
-- Name: contracts fk_kid; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.contracts
    ADD CONSTRAINT fk_kid FOREIGN KEY (kid_rg) REFERENCES public.kids(rg) ON DELETE CASCADE;


--
-- Name: contracts fk_responsible; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.contracts
    ADD CONSTRAINT fk_responsible FOREIGN KEY (responsible_cpf) REFERENCES public.responsibles(cpf) ON DELETE CASCADE;


--
-- Name: contracts fk_school; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.contracts
    ADD CONSTRAINT fk_school FOREIGN KEY (school_cnpj) REFERENCES public.schools(cnpj) ON DELETE CASCADE;


--
-- Name: invites invites_guester_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.invites
    ADD CONSTRAINT invites_guester_fkey FOREIGN KEY (guester) REFERENCES public.drivers(cnh);


--
-- Name: invites invites_requester_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.invites
    ADD CONSTRAINT invites_requester_fkey FOREIGN KEY (requester) REFERENCES public.schools(cnpj);


--
-- Name: kids kids_responsible_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.kids
    ADD CONSTRAINT kids_responsible_id_fkey FOREIGN KEY (responsible_id) REFERENCES public.responsibles(cpf) ON DELETE CASCADE;


--
-- Name: partners partners_driver_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.partners
    ADD CONSTRAINT partners_driver_id_fkey FOREIGN KEY (driver_cnh) REFERENCES public.drivers(cnh) ON DELETE CASCADE;


--
-- Name: partners partners_school_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.partners
    ADD CONSTRAINT partners_school_id_fkey FOREIGN KEY (school_cnpj) REFERENCES public.schools(cnpj) ON DELETE CASCADE;


--
-- Name: SCHEMA public; Type: ACL; Schema: -; Owner: postgres
--

REVOKE USAGE ON SCHEMA public FROM PUBLIC;
GRANT ALL ON SCHEMA public TO PUBLIC;


--
-- PostgreSQL database dump complete
--

