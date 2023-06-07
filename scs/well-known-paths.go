package scs

var wellKnownPaths = []string{
    "def", 
    "def/animated_model", 
    "def/camera", 
    "def/camera/city_start", 
    "def/camera/first_garage", 
    "def/camera/garage", 
    "def/camera/garage_poor", 
    "def/camera/garage_upgrade", 
    "def/camera/intro_garage", 
    "def/camera/purchase_truck", 
    "def/camera/ui", 
    "def/camera/units", 
    "def/cargo", 
    "def/cargo/acetylene", 
    "def/cargo/acid", 
    "def/cargo/ammonia", 
    "def/cargo/ammunition", 
    "def/cargo/apples_c", 
    "def/cargo/arsenic", 
    "def/cargo/beverages_c", 
    "def/cargo/big_bag_seed", 
    "def/cargo/brake_fluid", 
    "def/cargo/bricks", 
    "def/cargo/carbn_pwdr_c", 
    "def/cargo/carrots_c", 
    "def/cargo/cars_fr", 
    "def/cargo/car_balt1", 
    "def/cargo/car_balt2", 
    "def/cargo/car_d", 
    "def/cargo/car_f", 
    "def/cargo/car_ibe", 
    "def/cargo/car_it", 
    "def/cargo/chemicals", 
    "def/cargo/chem_sorb_c", 
    "def/cargo/chlorine", 
    "def/cargo/chlorine_t", 
    "def/cargo/clothes_c", 
    "def/cargo/concrete", 
    "def/cargo/concr_beams2", 
    "def/cargo/concr_cent", 
    "def/cargo/concr_stair", 
    "def/cargo/contamin", 
    "def/cargo/cyanide", 
    "def/cargo/diesel", 
    "def/cargo/digger1000", 
    "def/cargo/digger500", 
    "def/cargo/diggers", 
    "def/cargo/dynamite", 
    "def/cargo/emptytank", 
    "def/cargo/ethane", 
    "def/cargo/excavator", 
    "def/cargo/exhausts_c", 
    "def/cargo/explosives", 
    "def/cargo/fertilizer", 
    "def/cargo/fireworks", 
    "def/cargo/floorpanels", 
    "def/cargo/fluorine", 
    "def/cargo/forklifts", 
    "def/cargo/fueltanker", 
    "def/cargo/fuel_oil", 
    "def/cargo/glass", 
    "def/cargo/glass_packed", 
    "def/cargo/hchemicals", 
    "def/cargo/hipresstank", 
    "def/cargo/hmetal", 
    "def/cargo/hwaste", 
    "def/cargo/hydrochlor", 
    "def/cargo/hydrogen", 
    "def/cargo/iron_pipes", 
    "def/cargo/iveco_vans", 
    "def/cargo/kerosene", 
    "def/cargo/largetubes", 
    "def/cargo/lead", 
    "def/cargo/live_cattle", 
    "def/cargo/live_pigs", 
    "def/cargo/logs", 
    "def/cargo/lpg", 
    "def/cargo/lpg_t", 
    "def/cargo/lumber", 
    "def/cargo/lux_yacht", 
    "def/cargo/magnesium", 
    "def/cargo/marb_blck", 
    "def/cargo/marb_blck2", 
    "def/cargo/marb_slab", 
    "def/cargo/med_vaccine", 
    "def/cargo/mercuric", 
    "def/cargo/metal_beams", 
    "def/cargo/metal_pipes", 
    "def/cargo/mondeos", 
    "def/cargo/motor_oil_c", 
    "def/cargo/moto_tires", 
    "def/cargo/mtl_coil", 
    "def/cargo/neon", 
    "def/cargo/nitrocel", 
    "def/cargo/nitrogen", 
    "def/cargo/oil", 
    "def/cargo/oil_filt_c", 
    "def/cargo/olive_tree", 
    "def/cargo/outdr_flr_tl", 
    "def/cargo/overweight", 
    "def/cargo/pesticide", 
    "def/cargo/petrol", 
    "def/cargo/pet_food_c", 
    "def/cargo/phosphor", 
    "def/cargo/plant_substr", 
    "def/cargo/plast_film_c", 
    "def/cargo/plast_pipes", 
    "def/cargo/potahydro", 
    "def/cargo/potassium", 
    "def/cargo/propane", 
    "def/cargo/re_bars", 
    "def/cargo/rice_c", 
    "def/cargo/roof_tiles", 
    "def/cargo/salt_spice_c", 
    "def/cargo/scania_tr", 
    "def/cargo/sodchlor", 
    "def/cargo/sodhydro", 
    "def/cargo/sodium", 
    "def/cargo/sq_tub", 
    "def/cargo/straw_bales", 
    "def/cargo/sulfuric", 
    "def/cargo/tractors", 
    "def/cargo/train_part", 
    "def/cargo/train_part2", 
    "def/cargo/truck_batt_c", 
    "def/cargo/truck_rims_c", 
    "def/cargo/used_plast_c", 
    "def/cargo/vent_tube", 
    "def/cargo/vinegar_c", 
    "def/cargo/volvo_cars", 
    "def/cargo/volvo_tr", 
    "def/cargo/wallpanels", 
    "def/cargo/windml_eng", 
    "def/cargo/windml_tube", 
    "def/cargo/wooden_beams", 
    "def/cargo_obsolete", 
    "def/city", 
    "def/climate", 
    "def/climate/albedo", 
    "def/climate/arid", 
    "def/climate/black", 
    "def/climate/default", 
    "def/climate/desert", 
    "def/climate/grayscale", 
    "def/climate/integrity", 
    "def/climate/reference", 
    "def/company", 
    "def/company/acc", 
    "def/company/acc/in", 
    "def/company/acc/out", 
    "def/company/aci", 
    "def/company/aci/editor", 
    "def/company/aci/in", 
    "def/company/aci/out", 
    "def/company/ai", 
    "def/company/ai/editor", 
    "def/company/ai/in", 
    "def/company/ai/out", 
    "def/company/batisse_hs", 
    "def/company/batisse_hs/editor", 
    "def/company/batisse_hs/in", 
    "def/company/batisse_hs/out", 
    "def/company/bcp", 
    "def/company/bcp/editor", 
    "def/company/bcp/in", 
    "def/company/bcp/out", 
    "def/company/bhv", 
    "def/company/bhv/editor", 
    "def/company/bhv/in", 
    "def/company/bhv/out", 
    "def/company/blt_ru", 
    "def/company/blt_ru/in", 
    "def/company/blt_ru/out", 
    "def/company/blt_yacht_ru", 
    "def/company/blt_yacht_ru/in", 
    "def/company/blt_yacht_ru/out", 
    "def/company/cargotras", 
    "def/company/cargotras/editor", 
    "def/company/cargotras/in", 
    "def/company/cargotras/out", 
    "def/company/cemelt_fla", 
    "def/company/cemelt_fla/in", 
    "def/company/cemelt_fla/out", 
    "def/company/cemelt_fl_ru", 
    "def/company/cemelt_fl_ru/in", 
    "def/company/cemelt_fl_ru/out", 
    "def/company/cesare_smar", 
    "def/company/cesare_smar/editor", 
    "def/company/cesare_smar/in", 
    "def/company/cesare_smar/out", 
    "def/company/cgla", 
    "def/company/cgla/in", 
    "def/company/cgla/out", 
    "def/company/cnp", 
    "def/company/cnp/editor", 
    "def/company/cnp/in", 
    "def/company/cnp/out", 
    "def/company/comoto", 
    "def/company/comoto/in", 
    "def/company/comoto/out", 
    "def/company/cont_port", 
    "def/company/cont_port/in", 
    "def/company/cont_port/out", 
    "def/company/dfh", 
    "def/company/dfh/editor", 
    "def/company/dfh/in", 
    "def/company/dfh/out", 
    "def/company/domdepo", 
    "def/company/domdepo/editor", 
    "def/company/domdepo/in", 
    "def/company/domdepo/out", 
    "def/company/domdepo_ru", 
    "def/company/domdepo_ru/in", 
    "def/company/domdepo_ru/out", 
    "def/company/eolo_lines", 
    "def/company/eolo_lines/in", 
    "def/company/eolo_lines/out", 
    "def/company/euroacres", 
    "def/company/euroacres/editor", 
    "def/company/euroacres/in", 
    "def/company/euroacres/out", 
    "def/company/eurogoodies", 
    "def/company/eurogoodies/editor", 
    "def/company/eurogoodies/in", 
    "def/company/eurogoodies/out", 
    "def/company/exomar", 
    "def/company/exomar/in", 
    "def/company/exomar/out", 
    "def/company/fcp", 
    "def/company/fcp/editor", 
    "def/company/fcp/in", 
    "def/company/fcp/out", 
    "def/company/fle", 
    "def/company/fle/editor", 
    "def/company/fle/in", 
    "def/company/fle/out", 
    "def/company/gallia_ferry", 
    "def/company/gallia_ferry/in", 
    "def/company/gallia_ferry/out", 
    "def/company/gnt", 
    "def/company/gnt/in", 
    "def/company/gnt/out", 
    "def/company/han_expo", 
    "def/company/han_expo/editor", 
    "def/company/han_expo/in", 
    "def/company/han_expo/out", 
    "def/company/ika_bohag", 
    "def/company/ika_bohag/editor", 
    "def/company/ika_bohag/in", 
    "def/company/ika_bohag/out", 
    "def/company/itcc", 
    "def/company/itcc/editor", 
    "def/company/itcc/in", 
    "def/company/itcc/out", 
    "def/company/itcc_scrap", 
    "def/company/itcc_scrap/editor", 
    "def/company/itcc_scrap/in", 
    "def/company/itcc_scrap/out", 
    "def/company/kaarfor", 
    "def/company/kaarfor/editor", 
    "def/company/kaarfor/in", 
    "def/company/kaarfor/out", 
    "def/company/kathode", 
    "def/company/kathode/editor", 
    "def/company/kathode/in", 
    "def/company/kathode/out", 
    "def/company/kolico", 
    "def/company/kolico/editor", 
    "def/company/kolico/in", 
    "def/company/kolico/out", 
    "def/company/libellula", 
    "def/company/libellula/editor", 
    "def/company/libellula/in", 
    "def/company/libellula/out", 
    "def/company/lisette_log", 
    "def/company/lisette_log/editor", 
    "def/company/lisette_log/in", 
    "def/company/lisette_log/out", 
    "def/company/lkwlog", 
    "def/company/lkwlog/editor", 
    "def/company/lkwlog/in", 
    "def/company/lkwlog/out", 
    "def/company/marina", 
    "def/company/marina/in", 
    "def/company/marina/out", 
    "def/company/muromec", 
    "def/company/muromec/in", 
    "def/company/muromec/out", 
    "def/company/nbfc", 
    "def/company/nbfc/editor", 
    "def/company/nbfc/in", 
    "def/company/nbfc/out", 
    "def/company/nch_ru", 
    "def/company/nch_ru/in", 
    "def/company/nch_ru/out", 
    "def/company/nos_pat", 
    "def/company/nos_pat/in", 
    "def/company/nos_pat/out", 
    "def/company/ns_chem", 
    "def/company/ns_chem/in", 
    "def/company/ns_chem/out", 
    "def/company/ns_oil", 
    "def/company/ns_oil/editor", 
    "def/company/ns_oil/in", 
    "def/company/ns_oil/out", 
    "def/company/nucleon", 
    "def/company/nucleon/in", 
    "def/company/nucleon/out", 
    "def/company/ocean_sol", 
    "def/company/ocean_sol/in", 
    "def/company/ocean_sol/out", 
    "def/company/polarislines", 
    "def/company/polarislines/editor", 
    "def/company/polarislines/in", 
    "def/company/polarislines/out", 
    "def/company/posped", 
    "def/company/posped/editor", 
    "def/company/posped/in", 
    "def/company/posped/out", 
    "def/company/quarry", 
    "def/company/quarry/editor", 
    "def/company/quarry/in", 
    "def/company/quarry/out", 
    "def/company/radus", 
    "def/company/radus/in", 
    "def/company/radus/out", 
    "def/company/radus_ru", 
    "def/company/radus_ru/in", 
    "def/company/radus_ru/out", 
    "def/company/renar", 
    "def/company/renar/in", 
    "def/company/renar/out", 
    "def/company/renat", 
    "def/company/renat/editor", 
    "def/company/renat/in", 
    "def/company/renat/out", 
    "def/company/rt_log", 
    "def/company/rt_log/editor", 
    "def/company/rt_log/in", 
    "def/company/rt_log/out", 
    "def/company/sanbuilders", 
    "def/company/sanbuilders/editor", 
    "def/company/sanbuilders/in", 
    "def/company/sanbuilders/out", 
    "def/company/sanbuild_cem", 
    "def/company/sanbuild_cem/editor", 
    "def/company/sanbuild_cem/in", 
    "def/company/sanbuild_cem/out", 
    "def/company/sanbuild_hms", 
    "def/company/sanbuild_hms/editor", 
    "def/company/sanbuild_hms/in", 
    "def/company/sanbuild_hms/out", 
    "def/company/scania_dlr", 
    "def/company/scania_dlr/editor", 
    "def/company/scania_dlr/in", 
    "def/company/scs_paper", 
    "def/company/scs_paper/in", 
    "def/company/scs_paper/out", 
    "def/company/sellplan", 
    "def/company/sellplan/editor", 
    "def/company/sellplan/in", 
    "def/company/sellplan/out", 
    "def/company/sirin", 
    "def/company/sirin/in", 
    "def/company/sirin/out", 
    "def/company/skoda", 
    "def/company/skoda/editor", 
    "def/company/skoda/in", 
    "def/company/skoda/out", 
    "def/company/spinelli", 
    "def/company/spinelli/editor", 
    "def/company/spinelli/in", 
    "def/company/spinelli/out", 
    "def/company/steinbr_str", 
    "def/company/steinbr_str/editor", 
    "def/company/steinbr_str/in", 
    "def/company/steinbr_str/out", 
    "def/company/stokes", 
    "def/company/stokes/editor", 
    "def/company/stokes/in", 
    "def/company/stokes/out", 
    "def/company/subse", 
    "def/company/subse/in", 
    "def/company/subse/out", 
    "def/company/tesore_gust", 
    "def/company/tesore_gust/editor", 
    "def/company/tesore_gust/in", 
    "def/company/tesore_gust/out", 
    "def/company/tradeaux", 
    "def/company/tradeaux/editor", 
    "def/company/tradeaux/in", 
    "def/company/tradeaux/out", 
    "def/company/trainfoundry", 
    "def/company/trainfoundry/in", 
    "def/company/trainfoundry/out", 
    "def/company/trameri", 
    "def/company/trameri/editor", 
    "def/company/trameri/in", 
    "def/company/trameri/out", 
    "def/company/transinet", 
    "def/company/transinet/editor", 
    "def/company/transinet/in", 
    "def/company/transinet/out", 
    "def/company/tree_et", 
    "def/company/tree_et/editor", 
    "def/company/tree_et/in", 
    "def/company/tree_et/out", 
    "def/company/tree_et_log", 
    "def/company/tree_et_log/editor", 
    "def/company/tree_et_log/in", 
    "def/company/tree_et_log/out", 
    "def/company/vitas_pwr", 
    "def/company/vitas_pwr/in", 
    "def/company/vitas_pwr/out", 
    "def/company/voitureux", 
    "def/company/voitureux/editor", 
    "def/company/voitureux/in", 
    "def/company/voitureux/out", 
    "def/company/volvo_dlr", 
    "def/company/volvo_dlr/editor", 
    "def/company/volvo_dlr/in", 
    "def/company/wgcc", 
    "def/company/wgcc/editor", 
    "def/company/wgcc/in", 
    "def/company/wgcc/out", 
    "def/country", 
    "def/country/austria", 
    "def/country/belgium", 
    "def/country/bulgaria", 
    "def/country/czech", 
    "def/country/denmark", 
    "def/country/estonia", 
    "def/country/finland", 
    "def/country/france", 
    "def/country/germany", 
    "def/country/hungary", 
    "def/country/italy", 
    "def/country/latvia", 
    "def/country/lithuania", 
    "def/country/luxembourg", 
    "def/country/netherlands", 
    "def/country/norway", 
    "def/country/poland", 
    "def/country/portugal", 
    "def/country/romania", 
    "def/country/russia", 
    "def/country/slovakia", 
    "def/country/spain", 
    "def/country/sweden", 
    "def/country/switzerland", 
    "def/country/turkey", 
    "def/country/uk", 
    "def/desktop", 
    "def/economy_trailers_data", 
    "def/eye_tracking_presets", 
    "def/ferry", 
    "def/ferry/connection", 
    "def/initial_save", 
    "def/initial_save/normal", 
    "def/initial_save/preview", 
    "def/initial_save/test", 
    "def/jobs", 
    "def/photo_album", 
    "def/regional_settings", 
    "def/regional_settings/profiles", 
    "def/sign", 
    "def/sign/atlas", 
    "def/sign/board", 
    "def/sign/frame", 
    "def/sign/project", 
    "def/sign/template", 
    "def/stamp", 
    "def/tool", 
    "def/vehicle", 
    "def/vehicle/addon_hookups", 
    "def/vehicle/ai", 
    "def/vehicle/ai_wheel", 
    "def/vehicle/f_cover", 
    "def/vehicle/f_disc", 
    "def/vehicle/f_hub", 
    "def/vehicle/f_nuts", 
    "def/vehicle/f_rim", 
    "def/vehicle/f_tire", 
    "def/vehicle/f_wheel", 
    "def/vehicle/r_cover", 
    "def/vehicle/r_disc", 
    "def/vehicle/r_hub", 
    "def/vehicle/r_nuts", 
    "def/vehicle/r_rim", 
    "def/vehicle/r_tire", 
    "def/vehicle/r_wheel", 
    "def/vehicle/trailer", 
    "def/vehicle/trailer_cargo", 
    "def/vehicle/trailer_chain_types", 
    "def/vehicle/trailer_dealer", 
    "def/vehicle/trailer_defs", 
    "def/vehicle/trailer_desktop", 
    "def/vehicle/trailer_owned", 
    "def/vehicle/trailer_wheel", 
    "def/vehicle/truck", 
    "def/vehicle/truck/daf.2021", 
    "def/vehicle/truck/daf.2021/accessory", 
    "def/vehicle/truck/daf.2021/accessory/cup_holder", 
    "def/vehicle/truck/daf.2021/accessory/curtain_f", 
    "def/vehicle/truck/daf.2021/accessory/l_pillow", 
    "def/vehicle/truck/daf.2021/accessory/set_lglass", 
    "def/vehicle/truck/daf.2021/accessory/toyac", 
    "def/vehicle/truck/daf.2021/accessory/toybed", 
    "def/vehicle/truck/daf.2021/accessory/toybig", 
    "def/vehicle/truck/daf.2021/accessory/toyhang", 
    "def/vehicle/truck/daf.2021/accessory/toyseat", 
    "def/vehicle/truck/daf.2021/accessory/toystand", 
    "def/vehicle/truck/daf.xf", 
    "def/vehicle/truck/daf.xf/accessory", 
    "def/vehicle/truck/daf.xf/accessory/curtain_f", 
    "def/vehicle/truck/daf.xf/accessory/l_pillow", 
    "def/vehicle/truck/daf.xf/accessory/set_lglass", 
    "def/vehicle/truck/daf.xf/accessory/toyac", 
    "def/vehicle/truck/daf.xf/accessory/toybed", 
    "def/vehicle/truck/daf.xf/accessory/toybig", 
    "def/vehicle/truck/daf.xf/accessory/toyhang", 
    "def/vehicle/truck/daf.xf/accessory/toyseat", 
    "def/vehicle/truck/daf.xf/accessory/toystand", 
    "def/vehicle/truck/daf.xf_euro6", 
    "def/vehicle/truck/daf.xf_euro6/accessory", 
    "def/vehicle/truck/daf.xf_euro6/accessory/curtain_f", 
    "def/vehicle/truck/daf.xf_euro6/accessory/l_pillow", 
    "def/vehicle/truck/daf.xf_euro6/accessory/set_lglass", 
    "def/vehicle/truck/daf.xf_euro6/accessory/toyac", 
    "def/vehicle/truck/daf.xf_euro6/accessory/toybed", 
    "def/vehicle/truck/daf.xf_euro6/accessory/toybig", 
    "def/vehicle/truck/daf.xf_euro6/accessory/toyhang", 
    "def/vehicle/truck/daf.xf_euro6/accessory/toyseat", 
    "def/vehicle/truck/daf.xf_euro6/accessory/toystand", 
    "def/vehicle/truck/iveco.hiway", 
    "def/vehicle/truck/iveco.hiway/accessory", 
    "def/vehicle/truck/iveco.hiway/accessory/curtain_f", 
    "def/vehicle/truck/iveco.hiway/accessory/l_pillow", 
    "def/vehicle/truck/iveco.hiway/accessory/set_lglass", 
    "def/vehicle/truck/iveco.hiway/accessory/toyac", 
    "def/vehicle/truck/iveco.hiway/accessory/toybed", 
    "def/vehicle/truck/iveco.hiway/accessory/toybig", 
    "def/vehicle/truck/iveco.hiway/accessory/toyhang", 
    "def/vehicle/truck/iveco.hiway/accessory/toyseat", 
    "def/vehicle/truck/iveco.hiway/accessory/toystand", 
    "def/vehicle/truck/iveco.stralis", 
    "def/vehicle/truck/iveco.stralis/accessory", 
    "def/vehicle/truck/iveco.stralis/accessory/curtain_f", 
    "def/vehicle/truck/iveco.stralis/accessory/l_pillow", 
    "def/vehicle/truck/iveco.stralis/accessory/set_lglass", 
    "def/vehicle/truck/iveco.stralis/accessory/toyac", 
    "def/vehicle/truck/iveco.stralis/accessory/toybed", 
    "def/vehicle/truck/iveco.stralis/accessory/toybig", 
    "def/vehicle/truck/iveco.stralis/accessory/toyhang", 
    "def/vehicle/truck/iveco.stralis/accessory/toyseat", 
    "def/vehicle/truck/iveco.stralis/accessory/toystand", 
    "def/vehicle/truck/man.tgx", 
    "def/vehicle/truck/man.tgx/accessory", 
    "def/vehicle/truck/man.tgx/accessory/curtain_f", 
    "def/vehicle/truck/man.tgx/accessory/l_pillow", 
    "def/vehicle/truck/man.tgx/accessory/set_lglass", 
    "def/vehicle/truck/man.tgx/accessory/toyac", 
    "def/vehicle/truck/man.tgx/accessory/toybed", 
    "def/vehicle/truck/man.tgx/accessory/toybig", 
    "def/vehicle/truck/man.tgx/accessory/toyhang", 
    "def/vehicle/truck/man.tgx/accessory/toyseat", 
    "def/vehicle/truck/man.tgx/accessory/toystand", 
    "def/vehicle/truck/man.tgx_2020", 
    "def/vehicle/truck/man.tgx_2020/accessory", 
    "def/vehicle/truck/man.tgx_2020/accessory/cup_holder", 
    "def/vehicle/truck/man.tgx_2020/accessory/curtain_f", 
    "def/vehicle/truck/man.tgx_2020/accessory/l_pillow", 
    "def/vehicle/truck/man.tgx_2020/accessory/set_lglass", 
    "def/vehicle/truck/man.tgx_2020/accessory/toyac", 
    "def/vehicle/truck/man.tgx_2020/accessory/toybed", 
    "def/vehicle/truck/man.tgx_2020/accessory/toybig", 
    "def/vehicle/truck/man.tgx_2020/accessory/toyhang", 
    "def/vehicle/truck/man.tgx_2020/accessory/toyseat", 
    "def/vehicle/truck/man.tgx_2020/accessory/toystand", 
    "def/vehicle/truck/man.tgx_euro6", 
    "def/vehicle/truck/man.tgx_euro6/accessory", 
    "def/vehicle/truck/man.tgx_euro6/accessory/curtain_f", 
    "def/vehicle/truck/man.tgx_euro6/accessory/l_pillow", 
    "def/vehicle/truck/man.tgx_euro6/accessory/set_lglass", 
    "def/vehicle/truck/man.tgx_euro6/accessory/toyac", 
    "def/vehicle/truck/man.tgx_euro6/accessory/toybed", 
    "def/vehicle/truck/man.tgx_euro6/accessory/toybig", 
    "def/vehicle/truck/man.tgx_euro6/accessory/toyhang", 
    "def/vehicle/truck/man.tgx_euro6/accessory/toyseat", 
    "def/vehicle/truck/man.tgx_euro6/accessory/toystand", 
    "def/vehicle/truck/mercedes.actros", 
    "def/vehicle/truck/mercedes.actros/accessory", 
    "def/vehicle/truck/mercedes.actros/accessory/curtain_f", 
    "def/vehicle/truck/mercedes.actros/accessory/l_pillow", 
    "def/vehicle/truck/mercedes.actros/accessory/set_lglass", 
    "def/vehicle/truck/mercedes.actros/accessory/toyac", 
    "def/vehicle/truck/mercedes.actros/accessory/toybed", 
    "def/vehicle/truck/mercedes.actros/accessory/toybig", 
    "def/vehicle/truck/mercedes.actros/accessory/toyhang", 
    "def/vehicle/truck/mercedes.actros/accessory/toyseat", 
    "def/vehicle/truck/mercedes.actros/accessory/toystand", 
    "def/vehicle/truck/mercedes.actros2014", 
    "def/vehicle/truck/mercedes.actros2014/accessory", 
    "def/vehicle/truck/mercedes.actros2014/accessory/curtain_f", 
    "def/vehicle/truck/mercedes.actros2014/accessory/l_pillow", 
    "def/vehicle/truck/mercedes.actros2014/accessory/set_lglass", 
    "def/vehicle/truck/mercedes.actros2014/accessory/toyac", 
    "def/vehicle/truck/mercedes.actros2014/accessory/toybed", 
    "def/vehicle/truck/mercedes.actros2014/accessory/toybig", 
    "def/vehicle/truck/mercedes.actros2014/accessory/toyhang", 
    "def/vehicle/truck/mercedes.actros2014/accessory/toyseat", 
    "def/vehicle/truck/mercedes.actros2014/accessory/toystand", 
    "def/vehicle/truck/renault.magnum", 
    "def/vehicle/truck/renault.magnum/accessory", 
    "def/vehicle/truck/renault.magnum/accessory/curtain_f", 
    "def/vehicle/truck/renault.magnum/accessory/l_pillow", 
    "def/vehicle/truck/renault.magnum/accessory/set_lglass", 
    "def/vehicle/truck/renault.magnum/accessory/toyac", 
    "def/vehicle/truck/renault.magnum/accessory/toybed", 
    "def/vehicle/truck/renault.magnum/accessory/toybig", 
    "def/vehicle/truck/renault.magnum/accessory/toyhang", 
    "def/vehicle/truck/renault.magnum/accessory/toyseat", 
    "def/vehicle/truck/renault.magnum/accessory/toystand", 
    "def/vehicle/truck/renault.premium", 
    "def/vehicle/truck/renault.premium/accessory", 
    "def/vehicle/truck/renault.premium/accessory/curtain_f", 
    "def/vehicle/truck/renault.premium/accessory/l_pillow", 
    "def/vehicle/truck/renault.premium/accessory/set_lglass", 
    "def/vehicle/truck/renault.premium/accessory/toyac", 
    "def/vehicle/truck/renault.premium/accessory/toybed", 
    "def/vehicle/truck/renault.premium/accessory/toybig", 
    "def/vehicle/truck/renault.premium/accessory/toyhang", 
    "def/vehicle/truck/renault.premium/accessory/toyseat", 
    "def/vehicle/truck/renault.premium/accessory/toystand", 
    "def/vehicle/truck/renault.t", 
    "def/vehicle/truck/renault.t/accessory", 
    "def/vehicle/truck/renault.t/accessory/cup_holder", 
    "def/vehicle/truck/renault.t/accessory/curtain_f", 
    "def/vehicle/truck/renault.t/accessory/l_pillow", 
    "def/vehicle/truck/renault.t/accessory/set_lglass", 
    "def/vehicle/truck/renault.t/accessory/toyac", 
    "def/vehicle/truck/renault.t/accessory/toybed", 
    "def/vehicle/truck/renault.t/accessory/toybig", 
    "def/vehicle/truck/renault.t/accessory/toyhang", 
    "def/vehicle/truck/renault.t/accessory/toypanel", 
    "def/vehicle/truck/renault.t/accessory/toyseat", 
    "def/vehicle/truck/renault.t/accessory/toystand", 
    "def/vehicle/truck/scania.r", 
    "def/vehicle/truck/scania.r/accessory", 
    "def/vehicle/truck/scania.r/accessory/curtain_f", 
    "def/vehicle/truck/scania.r/accessory/l_pillow", 
    "def/vehicle/truck/scania.r/accessory/set_lglass", 
    "def/vehicle/truck/scania.r/accessory/toyac", 
    "def/vehicle/truck/scania.r/accessory/toybed", 
    "def/vehicle/truck/scania.r/accessory/toybig", 
    "def/vehicle/truck/scania.r/accessory/toyhang", 
    "def/vehicle/truck/scania.r/accessory/toyseat", 
    "def/vehicle/truck/scania.r/accessory/toystand", 
    "def/vehicle/truck/scania.r_2016", 
    "def/vehicle/truck/scania.r_2016/accessory", 
    "def/vehicle/truck/scania.r_2016/accessory/curtain_f", 
    "def/vehicle/truck/scania.r_2016/accessory/l_pillow", 
    "def/vehicle/truck/scania.r_2016/accessory/set_lglass", 
    "def/vehicle/truck/scania.r_2016/accessory/toyac", 
    "def/vehicle/truck/scania.r_2016/accessory/toybed", 
    "def/vehicle/truck/scania.r_2016/accessory/toybig", 
    "def/vehicle/truck/scania.r_2016/accessory/toyhang", 
    "def/vehicle/truck/scania.r_2016/accessory/toyseat", 
    "def/vehicle/truck/scania.r_2016/accessory/toystand", 
    "def/vehicle/truck/scania.streamline", 
    "def/vehicle/truck/scania.streamline/accessory", 
    "def/vehicle/truck/scania.streamline/accessory/curtain_f", 
    "def/vehicle/truck/scania.streamline/accessory/l_pillow", 
    "def/vehicle/truck/scania.streamline/accessory/set_lglass", 
    "def/vehicle/truck/scania.streamline/accessory/toyac", 
    "def/vehicle/truck/scania.streamline/accessory/toybed", 
    "def/vehicle/truck/scania.streamline/accessory/toybig", 
    "def/vehicle/truck/scania.streamline/accessory/toyhang", 
    "def/vehicle/truck/scania.streamline/accessory/toyseat", 
    "def/vehicle/truck/scania.streamline/accessory/toystand", 
    "def/vehicle/truck/scania.s_2016", 
    "def/vehicle/truck/scania.s_2016/accessory", 
    "def/vehicle/truck/scania.s_2016/accessory/curtain_f", 
    "def/vehicle/truck/scania.s_2016/accessory/l_pillow", 
    "def/vehicle/truck/scania.s_2016/accessory/set_lglass", 
    "def/vehicle/truck/scania.s_2016/accessory/toyac", 
    "def/vehicle/truck/scania.s_2016/accessory/toybed", 
    "def/vehicle/truck/scania.s_2016/accessory/toybig", 
    "def/vehicle/truck/scania.s_2016/accessory/toyhang", 
    "def/vehicle/truck/scania.s_2016/accessory/toyseat", 
    "def/vehicle/truck/scania.s_2016/accessory/toystand", 
    "def/vehicle/truck/volvo.fh16", 
    "def/vehicle/truck/volvo.fh16/accessory", 
    "def/vehicle/truck/volvo.fh16/accessory/curtain_f", 
    "def/vehicle/truck/volvo.fh16/accessory/l_pillow", 
    "def/vehicle/truck/volvo.fh16/accessory/set_lglass", 
    "def/vehicle/truck/volvo.fh16/accessory/toyac", 
    "def/vehicle/truck/volvo.fh16/accessory/toybed", 
    "def/vehicle/truck/volvo.fh16/accessory/toybig", 
    "def/vehicle/truck/volvo.fh16/accessory/toyhang", 
    "def/vehicle/truck/volvo.fh16/accessory/toyseat", 
    "def/vehicle/truck/volvo.fh16/accessory/toystand", 
    "def/vehicle/truck/volvo.fh16_2012", 
    "def/vehicle/truck/volvo.fh16_2012/accessory", 
    "def/vehicle/truck/volvo.fh16_2012/accessory/curtain_f", 
    "def/vehicle/truck/volvo.fh16_2012/accessory/l_pillow", 
    "def/vehicle/truck/volvo.fh16_2012/accessory/set_lglass", 
    "def/vehicle/truck/volvo.fh16_2012/accessory/toyac", 
    "def/vehicle/truck/volvo.fh16_2012/accessory/toybed", 
    "def/vehicle/truck/volvo.fh16_2012/accessory/toybig", 
    "def/vehicle/truck/volvo.fh16_2012/accessory/toyhang", 
    "def/vehicle/truck/volvo.fh16_2012/accessory/toyseat", 
    "def/vehicle/truck/volvo.fh16_2012/accessory/toystand", 
    "def/vehicle/truck_company", 
    "def/vehicle/truck_company_uk", 
    "def/vehicle/truck_dealer", 
    "def/vehicle/truck_dealer_uk", 
    "def/vehicle/truck_desktop", 
    "def/vehicle/truck_desktop_uk", 
    "def/vehicle/t_wheel", 
    "def/world", 
    "material", 
    "material/building", 
    "material/colors", 
    "material/concrete", 
    "material/custom", 
    "material/decal", 
    "material/deferred", 
    "material/editor", 
    "material/environment", 
    "material/flags", 
    "material/ground", 
    "material/legacy", 
    "material/logo", 
    "material/metal", 
    "material/obsolete", 
    "material/overlay", 
    "material/photo_album", 
    "material/prefab", 
    "material/road", 
    "material/rocks", 
    "material/ropes", 
    "material/sign", 
    "material/special", 
    "material/subset", 
    "material/terrain", 
    "material/tiled", 
    "material/ui", 
    "material/ui/accessory", 
    "material/ui/accessory/truck", 
    "material/ui/accessory/truck/upgrade", 
    "material/ui/accessory/truck/upgrade/interior_decors", 
    "material/ui/accessory/truck/upgrade/interior_decors/toypanel", 
    "material/ui/accessory/truck/upgrade/interior_decors/toypanel/cellphone", 
    "material/ui/accessory/truck/upgrade/interior_decors/toypanel/tablet", 
    "material/ui/packages", 
    "ui", 
    "ui/company_manager", 
    "ui/dashboard", 
    "ui/desc", 
    "ui/editor", 
    "ui/layout", 
    "ui/license_plate", 
    "ui/multiplayer", 
    "ui/online", 
    "ui/options", 
    "ui/player_tag", 
    "ui/template", 
    "vehicle", 
    "vehicle/ai", 
    "vehicle/car", 
    "vehicle/driver", 
    "vehicle/share", 
    "vehicle/trailer_eu", 
    "vehicle/trailer_owned", 
    "vehicle/truck", 
    "vehicle/truck/upgrade", 
    "vehicle/truck/upgrade/interior_decors", 
    "vehicle/truck/upgrade/interior_decors/curtains", 
    "vehicle/truck/upgrade/interior_decors/curtains/daf_2021", 
    "vehicle/truck/upgrade/interior_decors/curtains/daf_xf", 
    "vehicle/truck/upgrade/interior_decors/curtains/daf_xf_euro6", 
    "vehicle/truck/upgrade/interior_decors/curtains/iveco_hiway", 
    "vehicle/truck/upgrade/interior_decors/curtains/iveco_stralis", 
    "vehicle/truck/upgrade/interior_decors/curtains/man_tgx", 
    "vehicle/truck/upgrade/interior_decors/curtains/man_tgx_euro6", 
    "vehicle/truck/upgrade/interior_decors/curtains/mercedes_actros_2009", 
    "vehicle/truck/upgrade/interior_decors/curtains/mercedes_actros_2014", 
    "vehicle/truck/upgrade/interior_decors/curtains/renault_magnum_2009", 
    "vehicle/truck/upgrade/interior_decors/curtains/renault_premium", 
    "vehicle/truck/upgrade/interior_decors/curtains/renault_t", 
    "vehicle/truck/upgrade/interior_decors/curtains/scania_2016", 
    "vehicle/truck/upgrade/interior_decors/curtains/scania_rcab_2009", 
    "vehicle/truck/upgrade/interior_decors/curtains/scania_streamline", 
    "vehicle/truck/upgrade/interior_decors/curtains/volvo_fh16_2009", 
    "vehicle/truck/upgrade/interior_decors/curtains/volvo_fh16_2012", 
    "vehicle/truck/upgrade/interior_decors/gps", 
    "vehicle/truck/upgrade/interior_decors/toyac", 
    "vehicle/truck/upgrade/interior_decors/toybed", 
    "vehicle/truck/upgrade/interior_decors/toybig", 
    "vehicle/truck/upgrade/interior_decors/toyglass", 
    "vehicle/truck/upgrade/interior_decors/toyhang", 
    "vehicle/truck/upgrade/interior_decors/toypanel", 
    "vehicle/truck/upgrade/interior_decors/toypanel/cellphone", 
    "vehicle/truck/upgrade/interior_decors/toypanel/tablet", 
    "vehicle/truck/upgrade/interior_decors/toyseat", 
    "vehicle/truck/upgrade/interior_decors/toystand", 
    "vehicle/wheel",
}