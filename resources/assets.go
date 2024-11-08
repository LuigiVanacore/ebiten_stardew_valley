package resources

import (
	_ "embed"
)
var (
	//go:embed graphics\character\down\0.png
	Character_Down_0 []byte

	//go:embed graphics\character\down\1.png
	Character_Down_1 []byte

	//go:embed graphics\character\down\2.png
	Character_Down_2 []byte

	//go:embed graphics\character\down\3.png
	Character_Down_3 []byte

	//go:embed graphics\character\down_axe\0.png
	Character_Down_Axe_0 []byte

	//go:embed graphics\character\down_axe\1.png
	Character_Down_Axe_1 []byte

	//go:embed graphics\character\down_hoe\0.png
	Character_Down_Hoe_0 []byte

	//go:embed graphics\character\down_hoe\1.png
	Character_Down_Hoe_1 []byte

	//go:embed graphics\character\down_idle\0.png
	Character_Down_Idle_0 []byte

	//go:embed graphics\character\down_idle\1.png
	Character_Down_Idle_1 []byte

	//go:embed graphics\character\down_water\0.png
	Character_Down_Water_0 []byte

	//go:embed graphics\character\down_water\1.png
	Character_Down_Water_1 []byte

	//go:embed graphics\character\left\0.png
	Character_Left_0 []byte

	//go:embed graphics\character\left\1.png
	Character_Left_1 []byte

	//go:embed graphics\character\left\2.png
	Character_Left_2 []byte

	//go:embed graphics\character\left\3.png
	Character_Left_3 []byte

	//go:embed graphics\character\left_axe\0.png
	Character_Left_Axe_0 []byte

	//go:embed graphics\character\left_axe\1.png
	Character_Left_Axe_1 []byte

	//go:embed graphics\character\left_hoe\0.png
	Character_Left_Hoe_0 []byte

	//go:embed graphics\character\left_hoe\1.png
	Character_Left_Hoe_1 []byte

	//go:embed graphics\character\left_idle\0.png
	Character_Left_Idle_0 []byte

	//go:embed graphics\character\left_idle\1.png
	Character_Left_Idle_1 []byte

	//go:embed graphics\character\left_water\0.png
	Character_Left_Water_0 []byte

	//go:embed graphics\character\left_water\1.png
	Character_Left_Water_1 []byte

	//go:embed graphics\character\right\0.png
	Character_Right_0 []byte

	//go:embed graphics\character\right\1.png
	Character_Right_1 []byte

	//go:embed graphics\character\right\2.png
	Character_Right_2 []byte

	//go:embed graphics\character\right\3.png
	Character_Right_3 []byte

	//go:embed graphics\character\right_axe\0.png
	Character_Right_Axe_0 []byte

	//go:embed graphics\character\right_axe\1.png
	Character_Right_Axe_1 []byte

	//go:embed graphics\character\right_hoe\0.png
	Character_Right_Hoe_0 []byte

	//go:embed graphics\character\right_hoe\1.png
	Character_Right_Hoe_1 []byte

	//go:embed graphics\character\right_idle\0.png
	Character_Right_Idle_0 []byte

	//go:embed graphics\character\right_idle\1.png
	Character_Right_Idle_1 []byte

	//go:embed graphics\character\right_water\0.png
	Character_Right_Water_0 []byte

	//go:embed graphics\character\right_water\1.png
	Character_Right_Water_1 []byte

	//go:embed graphics\character\up\0.png
	Character_Up_0 []byte

	//go:embed graphics\character\up\1.png
	Character_Up_1 []byte

	//go:embed graphics\character\up\2.png
	Character_Up_2 []byte

	//go:embed graphics\character\up\3.png
	Character_Up_3 []byte

	//go:embed graphics\character\up_axe\0.png
	Character_Up_Axe_0 []byte

	//go:embed graphics\character\up_axe\1.png
	Character_Up_Axe_1 []byte

	//go:embed graphics\character\up_hoe\0.png
	Character_Up_Hoe_0 []byte

	//go:embed graphics\character\up_hoe\1.png
	Character_Up_Hoe_1 []byte

	//go:embed graphics\character\up_idle\0.png
	Character_Up_Idle_0 []byte

	//go:embed graphics\character\up_idle\1.png
	Character_Up_Idle_1 []byte

	//go:embed graphics\character\up_water\0.png
	Character_Up_Water_0 []byte

	//go:embed graphics\character\up_water\1.png
	Character_Up_Water_1 []byte

	//go:embed graphics\environment\Bridge.png
	Environment_Bridge []byte

	//go:embed graphics\environment\Collision.png
	Environment_Collision []byte

	//go:embed graphics\environment\Fences.png
	Environment_Fences []byte

	//go:embed graphics\environment\Grass.png
	Environment_Grass []byte

	//go:embed graphics\environment\Hills.png
	Environment_Hills []byte

	//go:embed graphics\environment\HouseDecoration.png
	Environment_House_Decoration []byte

	//go:embed graphics\environment\House.png
	Environment_House []byte

	//go:embed graphics\environment\Paths.png
	Environment_Paths []byte

	//go:embed graphics\environment\PlantDecoration.png
	Environment_Plant_Decoration []byte

	//go:embed graphics\environment\WaterDecoration.png
	Environment_Water_Decoration []byte

	//go:embed graphics\environment\Water.png
	Environment_Water []byte

	//go:embed graphics\environment\interaction.png
	Environment_Interaction []byte

	//go:embed graphics\fruit\apple.png
	Fruit_Apple []byte

	//go:embed graphics\fruit\corn\0.png
	Fruit_Corn_0 []byte

	//go:embed graphics\fruit\corn\1.png
	Fruit_Corn_1 []byte

	//go:embed graphics\fruit\corn\2.png
	Fruit_Corn_2 []byte

	//go:embed graphics\fruit\corn\3.png
	Fruit_Corn_3 []byte

	//go:embed graphics\fruit\tomato\0.png
	Fruit_Tomato_0 []byte

	//go:embed graphics\fruit\tomato\1.png
	Fruit_Tomato_1 []byte

	//go:embed graphics\fruit\tomato\2.png
	Fruit_Tomato_2 []byte

	//go:embed graphics\fruit\tomato\3.png
	Fruit_Tomato_3 []byte

	//go:embed graphics\objects\bush.png
	Objects_Bush []byte

	//go:embed graphics\objects\flower.png
	Objects_Flower []byte

	//go:embed graphics\objects\merchant.png
	Objects_Merchant []byte

	//go:embed graphics\objects\mushroom.png
	Objects_Mushroom []byte

	//go:embed graphics\objects\mushrooms.png
	Objects_Mushrooms []byte

	//go:embed graphics\objects\stump_medium.png
	Objects_Stump_Medium []byte

	//go:embed graphics\objects\stump_small.png
	Objects_Stump_Small []byte

	//go:embed graphics\objects\sunflower.png
	Objects_Sunflower []byte

	//go:embed graphics\objects\tree_medium.png
	Objects_Tree_Medium []byte

	//go:embed graphics\objects\tree_small.png
	Objects_Tree_Small []byte

	//go:embed graphics\overlay\axe.png
	Overlay_Axe []byte

	//go:embed graphics\overlay\corn.png
	Overlay_Corn []byte

	//go:embed graphics\overlay\hoe.png
	Overlay_Hoe []byte

	//go:embed graphics\overlay\tomato.png
	Overlay_Tomato []byte

	//go:embed graphics\overlay\water.png
	Overlay_Water []byte

	//go:embed graphics\rain\drops\0.png
	Rain_Drops_0 []byte

	//go:embed graphics\rain\drops\1.png
	Rain_Drops_1 []byte

	//go:embed graphics\rain\drops\2.png
	Rain_Drops_2 []byte

	//go:embed graphics\rain\floor\0.png
	Rain_Floor_0 []byte

	//go:embed graphics\rain\floor\1.png
	Rain_Floor_1 []byte

	//go:embed graphics\rain\floor\2.png
	Rain_Floor_2 []byte

	//go:embed graphics\soil\b.png
	Soil_B []byte

	//go:embed graphics\soil\bl.png
	Soil_Bl []byte

	//go:embed graphics\soil\bm.png
	Soil_Bm []byte

	//go:embed graphics\soil\br.png
	Soil_Br []byte

	//go:embed graphics\soil\l.png
	Soil_L []byte

	//go:embed graphics\soil\lm.png
	Soil_Lm []byte

	//go:embed graphics\soil\lr.png
	Soil_Lr []byte

	//go:embed graphics\soil\lrb.png
	Soil_Lrb []byte

	//go:embed graphics\soil\lrt.png
	Soil_Lrt []byte

	//go:embed graphics\soil\o.png
	Soil_O []byte

	//go:embed graphics\soil\r.png
	Soil_R []byte

	//go:embed graphics\soil\rm.png
	Soil_Rm []byte

	//go:embed graphics\soil\soil.png
	Soil_Soil []byte

	//go:embed graphics\soil\t.png
	Soil_T []byte

	//go:embed graphics\soil\tb.png
	Soil_Tb []byte

	//go:embed graphics\soil\tbl.png
	Soil_Tbl []byte

	//go:embed graphics\soil\tbr.png
	Soil_Tbr []byte

	//go:embed graphics\soil\tl.png
	Soil_Tl []byte

	//go:embed graphics\soil\tm.png
	Soil_Tm []byte

	//go:embed graphics\soil\tr.png
	Soil_Tr []byte

	//go:embed graphics\soil\x.png
	Soil_X []byte

	//go:embed graphics\soil_water\0.png
	Soil_Water_0 []byte

	//go:embed graphics\soil_water\1.png
	Soil_Water_1 []byte

	//go:embed graphics\soil_water\2.png
	Soil_Water_2 []byte

	//go:embed graphics\stumps\large.png
	Stumps_Large []byte

	//go:embed graphics\stumps\small.png
	Stumps_Small []byte

	//go:embed graphics\water\0.png
	Water_0 []byte

	//go:embed graphics\water\1.png
	Water_1 []byte

	//go:embed graphics\water\2.png
	Water_2 []byte

	//go:embed graphics\water\3.png
	Water_3 []byte

	//go:embed graphics\world\ground.png
	World_Ground []byte
)
