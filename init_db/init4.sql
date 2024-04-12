CREATE OR REPLACE FUNCTION createBannerWithTags(
    tag_ids int[],
    feature_id int,
    title varchar(255),
    text text,
    url text,
    is_active boolean
)
RETURNS int
language plpgsql
as $$
DECLARE 
    id_banner int;
begin
	
	id_banner := createBanner(
    feature_id,
    title,
    text,
    url,
    is_active 
    );
	
	PERFORM createConnectionBannerTags(id_banner,tag_ids);
    
    return id_banner;
end;
$$;

CREATE OR REPLACE FUNCTION updateBannerWithTags(
    id_banner_in int,
    tag_ids_in int[],
    feature_id_in int,
    title_in varchar(255),
    text_in text,
    url_in text,
    is_active_in boolean)
RETURNS void
language plpgsql
as $$
begin
	
    PERFORM deleteConnectionBannerTags(id_banner_in);

	UPDATE banners
    set title = title_in, text = text_in, url = url_in, feature_id = feature_id_in, visible = is_active_in, update_time = NOW()
	WHERE id = id_banner_in;

	PERFORM createConnectionBannerTags(id_banner_in,tag_ids_in);

end;
$$;