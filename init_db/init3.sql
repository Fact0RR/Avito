CREATE OR REPLACE FUNCTION createBanner(
    feature_id int,
    title varchar(255),
    text text,
    url text,
    is_active boolean
)
RETURNS int
language plpgsql
as $$
DECLARE id_banner int;
begin
    insert into banners (title, text, url, feature_id, visible, create_time, update_time) values(
        title,
        text,
        url,
        feature_id,
        is_active,
        NOW(),
        NOW()
    ) RETURNING id into id_banner;
    return id_banner;
end;
$$;

CREATE OR REPLACE FUNCTION createConnectionBannerTags(
    id_banner int,
    tag_ids int[]
)
RETURNS void
language plpgsql
as $$
DECLARE 
    _elem int;
begin
    FOREACH _elem IN ARRAY tag_ids
    LOOP 

        insert into b_t (banner_id, tag_id) values
        (id_banner, _elem);

    END LOOP;
    return;
end;
$$;

CREATE OR REPLACE FUNCTION deleteConnectionBannerTags(
    id_banner int
)
RETURNS void
language plpgsql
as $$
begin
    DELETE FROM B_T
    WHERE banner_id = id_banner;
    return;
end;
$$;
