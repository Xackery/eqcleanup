package paw

import (
	"github.com/xackery/goeq/loot"
	//"database/sql"
)

var loottableentries []loot.LootTableEntries = []loot.LootTableEntries{
	{Loottable_id: 2971, Lootdrop_id: 5693, Multiplier: 1, Droplimit: 1, Mindrop: 0, Probability: 100.000000},
	{Loottable_id: 11627, Lootdrop_id: 18843, Multiplier: 1, Droplimit: 0, Mindrop: 0, Probability: 100.000000},
	{Loottable_id: 11627, Lootdrop_id: 100511, Multiplier: 1, Droplimit: 1, Mindrop: 0, Probability: 100.000000},
	{Loottable_id: 11627, Lootdrop_id: 106094, Multiplier: 1, Droplimit: 0, Mindrop: 0, Probability: 100.000000},
	{Loottable_id: 11627, Lootdrop_id: 106095, Multiplier: 1, Droplimit: 1, Mindrop: 0, Probability: 100.000000},
	{Mindrop: 0, Probability: 100.000000, Loottable_id: 11843, Lootdrop_id: 18288, Multiplier: 1, Droplimit: 1},
	{Loottable_id: 11889, Lootdrop_id: 19188, Multiplier: 4, Droplimit: 1, Mindrop: 0, Probability: 55.000000},
	{Loottable_id: 90208, Lootdrop_id: 90331, Multiplier: 1, Droplimit: 1, Mindrop: 0, Probability: 100.000000},
	{Loottable_id: 90208, Lootdrop_id: 105761, Multiplier: 1, Droplimit: 1, Mindrop: 1, Probability: 100.000000},
	{Loottable_id: 90208, Lootdrop_id: 105877, Multiplier: 1, Droplimit: 1, Mindrop: 0, Probability: 100.000000},
	{Loottable_id: 91623, Lootdrop_id: 96053, Multiplier: 1, Droplimit: 1, Mindrop: 0, Probability: 100.000000},
	{Lootdrop_id: 96054, Multiplier: 1, Droplimit: 1, Mindrop: 0, Probability: 100.000000, Loottable_id: 91623},
	{Mindrop: 0, Probability: 100.000000, Loottable_id: 91631, Lootdrop_id: 96165, Multiplier: 1, Droplimit: 1},
	{Lootdrop_id: 96166, Multiplier: 1, Droplimit: 1, Mindrop: 0, Probability: 100.000000, Loottable_id: 91631},
	{Mindrop: 0, Probability: 100.000000, Loottable_id: 92466, Lootdrop_id: 102478, Multiplier: 1, Droplimit: 1},
	{Loottable_id: 92466, Lootdrop_id: 102479, Multiplier: 1, Droplimit: 1, Mindrop: 1, Probability: 100.000000},
	{Multiplier: 1, Droplimit: 0, Mindrop: 0, Probability: 100.000000, Loottable_id: 92466, Lootdrop_id: 105699},
	{Loottable_id: 92466, Lootdrop_id: 105877, Multiplier: 1, Droplimit: 1, Mindrop: 0, Probability: 100.000000},
	{Mindrop: 0, Probability: 100.000000, Loottable_id: 92467, Lootdrop_id: 18843, Multiplier: 1, Droplimit: 0},
	{Loottable_id: 92467, Lootdrop_id: 100511, Multiplier: 1, Droplimit: 1, Mindrop: 0, Probability: 100.000000},
	{Mindrop: 0, Probability: 100.000000, Loottable_id: 92467, Lootdrop_id: 106095, Multiplier: 1, Droplimit: 1},
	{Lootdrop_id: 106098, Multiplier: 1, Droplimit: 0, Mindrop: 0, Probability: 100.000000, Loottable_id: 92467},
	{Droplimit: 0, Mindrop: 0, Probability: 100.000000, Loottable_id: 92468, Lootdrop_id: 102483, Multiplier: 1},
	{Loottable_id: 92469, Lootdrop_id: 102485, Multiplier: 1, Droplimit: 0, Mindrop: 0, Probability: 100.000000},
	{Loottable_id: 92469, Lootdrop_id: 102486, Multiplier: 1, Droplimit: 0, Mindrop: 0, Probability: 100.000000},
	{Loottable_id: 92469, Lootdrop_id: 105877, Multiplier: 1, Droplimit: 1, Mindrop: 0, Probability: 100.000000},
	{Droplimit: 1, Mindrop: 0, Probability: 100.000000, Loottable_id: 92469, Lootdrop_id: 106092, Multiplier: 1},
	{Droplimit: 0, Mindrop: 0, Probability: 100.000000, Loottable_id: 92470, Lootdrop_id: 102488, Multiplier: 1},
	{Probability: 100.000000, Loottable_id: 92470, Lootdrop_id: 102489, Multiplier: 1, Droplimit: 0, Mindrop: 0},
	{Probability: 100.000000, Loottable_id: 92470, Lootdrop_id: 105869, Multiplier: 1, Droplimit: 1, Mindrop: 0},
	{Probability: 100.000000, Loottable_id: 92470, Lootdrop_id: 106093, Multiplier: 1, Droplimit: 1, Mindrop: 0},
	{Probability: 100.000000, Loottable_id: 92471, Lootdrop_id: 102491, Multiplier: 1, Droplimit: 0, Mindrop: 0},
	{Multiplier: 1, Droplimit: 0, Mindrop: 0, Probability: 100.000000, Loottable_id: 92471, Lootdrop_id: 102492},
	{Loottable_id: 92471, Lootdrop_id: 106091, Multiplier: 1, Droplimit: 1, Mindrop: 0, Probability: 100.000000},
	{Loottable_id: 92472, Lootdrop_id: 102494, Multiplier: 1, Droplimit: 0, Mindrop: 0, Probability: 100.000000},
	{Lootdrop_id: 102495, Multiplier: 1, Droplimit: 0, Mindrop: 0, Probability: 100.000000, Loottable_id: 92472},
	{Mindrop: 0, Probability: 100.000000, Loottable_id: 92472, Lootdrop_id: 105877, Multiplier: 1, Droplimit: 1},
	{Loottable_id: 92472, Lootdrop_id: 106090, Multiplier: 1, Droplimit: 1, Mindrop: 0, Probability: 100.000000},
	{Probability: 100.000000, Loottable_id: 93724, Lootdrop_id: 100511, Multiplier: 1, Droplimit: 1, Mindrop: 0},
	{Droplimit: 1, Mindrop: 0, Probability: 100.000000, Loottable_id: 93724, Lootdrop_id: 106095, Multiplier: 1},
	{Probability: 100.000000, Loottable_id: 93724, Lootdrop_id: 106097, Multiplier: 1, Droplimit: 0, Mindrop: 0},
	{Lootdrop_id: 106098, Multiplier: 1, Droplimit: 0, Mindrop: 0, Probability: 100.000000, Loottable_id: 93724},
	{Probability: 100.000000, Loottable_id: 93725, Lootdrop_id: 105875, Multiplier: 1, Droplimit: 1, Mindrop: 0},
	{Droplimit: 1, Mindrop: 0, Probability: 100.000000, Loottable_id: 93725, Lootdrop_id: 106099, Multiplier: 1},
	{Mindrop: 0, Probability: 100.000000, Loottable_id: 93725, Lootdrop_id: 106100, Multiplier: 1, Droplimit: 0},
	{Lootdrop_id: 105875, Multiplier: 1, Droplimit: 1, Mindrop: 0, Probability: 100.000000, Loottable_id: 93726},
	{Droplimit: 1, Mindrop: 0, Probability: 100.000000, Loottable_id: 93726, Lootdrop_id: 106101, Multiplier: 1},
	{Mindrop: 0, Probability: 100.000000, Loottable_id: 93726, Lootdrop_id: 106102, Multiplier: 1, Droplimit: 0},
	{Loottable_id: 93727, Lootdrop_id: 18843, Multiplier: 1, Droplimit: 0, Mindrop: 0, Probability: 100.000000},
	{Lootdrop_id: 100511, Multiplier: 1, Droplimit: 1, Mindrop: 0, Probability: 100.000000, Loottable_id: 93727},
	{Droplimit: 0, Mindrop: 0, Probability: 100.000000, Loottable_id: 93727, Lootdrop_id: 106094, Multiplier: 1},
	{Droplimit: 1, Mindrop: 0, Probability: 100.000000, Loottable_id: 93727, Lootdrop_id: 106095, Multiplier: 1},
	{Mindrop: 0, Probability: 100.000000, Loottable_id: 93728, Lootdrop_id: 18843, Multiplier: 1, Droplimit: 0},
	{Loottable_id: 93728, Lootdrop_id: 105875, Multiplier: 1, Droplimit: 1, Mindrop: 0, Probability: 100.000000},
	{Loottable_id: 93728, Lootdrop_id: 106094, Multiplier: 1, Droplimit: 0, Mindrop: 0, Probability: 100.000000},
	{Droplimit: 1, Mindrop: 0, Probability: 100.000000, Loottable_id: 93728, Lootdrop_id: 106095, Multiplier: 1},
	{Droplimit: 0, Mindrop: 0, Probability: 100.000000, Loottable_id: 93729, Lootdrop_id: 18843, Multiplier: 1},
	{Probability: 100.000000, Loottable_id: 93729, Lootdrop_id: 105875, Multiplier: 1, Droplimit: 1, Mindrop: 0},
	{Multiplier: 1, Droplimit: 1, Mindrop: 0, Probability: 100.000000, Loottable_id: 93729, Lootdrop_id: 106095},
	{Loottable_id: 93729, Lootdrop_id: 106098, Multiplier: 1, Droplimit: 0, Mindrop: 0, Probability: 100.000000},
	{Loottable_id: 93730, Lootdrop_id: 18843, Multiplier: 1, Droplimit: 0, Mindrop: 0, Probability: 100.000000},
	{Mindrop: 0, Probability: 100.000000, Loottable_id: 93730, Lootdrop_id: 105875, Multiplier: 1, Droplimit: 1},
	{Loottable_id: 93730, Lootdrop_id: 106094, Multiplier: 1, Droplimit: 0, Mindrop: 0, Probability: 100.000000},
	{Mindrop: 0, Probability: 100.000000, Loottable_id: 93730, Lootdrop_id: 106095, Multiplier: 1, Droplimit: 1},
	{Loottable_id: 93731, Lootdrop_id: 18843, Multiplier: 1, Droplimit: 0, Mindrop: 0, Probability: 100.000000},
	{Loottable_id: 93731, Lootdrop_id: 105873, Multiplier: 1, Droplimit: 1, Mindrop: 0, Probability: 100.000000},
	{Lootdrop_id: 106094, Multiplier: 1, Droplimit: 0, Mindrop: 0, Probability: 100.000000, Loottable_id: 93731},
	{Probability: 100.000000, Loottable_id: 93731, Lootdrop_id: 106112, Multiplier: 1, Droplimit: 1, Mindrop: 0},
	{Droplimit: 0, Mindrop: 0, Probability: 100.000000, Loottable_id: 93732, Lootdrop_id: 18843, Multiplier: 1},
	{Multiplier: 1, Droplimit: 1, Mindrop: 0, Probability: 100.000000, Loottable_id: 93732, Lootdrop_id: 105871},
	{Lootdrop_id: 106094, Multiplier: 1, Droplimit: 0, Mindrop: 0, Probability: 100.000000, Loottable_id: 93732},
	{Droplimit: 1, Mindrop: 0, Probability: 100.000000, Loottable_id: 93732, Lootdrop_id: 106112, Multiplier: 1},
}