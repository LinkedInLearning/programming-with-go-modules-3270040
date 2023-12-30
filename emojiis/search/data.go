package search

// Emoji represents an Emoji with description and tag data
type Emoji struct {
	Emoji    string
	Label    string
	Tags     []string
	Emoticon string
}

// emojis is a super small subset of emoji data (derived from https://github.com/milesj/emojibase)
var emojis = []Emoji{
	{Label: "grinning face", Tags: []string{"face", "grin"}, Emoji: "😀"},
	{Label: "grinning face with big eyes", Tags: []string{"face", "mouth", "open", "smile"}, Emoji: "😃"},
	{Label: "grinning face with smiling eyes", Tags: []string{"eye", "face", "mouth", "open", "smile"}, Emoji: "😄", Emoticon: ":D"},
	{Label: "beaming face with smiling eyes", Tags: []string{"eye", "face", "grin", "smile"}, Emoji: "😁"},
	{Label: "grinning squinting face", Tags: []string{"face", "laugh", "mouth", "satisfied", "smile"}, Emoji: "😆", Emoticon: "xD"},
	{Label: "grinning face with sweat", Tags: []string{"cold", "face", "open", "smile", "sweat"}, Emoji: "😅"},
	{Label: "rolling on the floor laughing", Tags: []string{"face", "floor", "laugh", "rofl", "rolling", "rotfl"}, Emoji: "🤣", Emoticon: ":'D"},
	{Label: "face with tears of joy", Tags: []string{"face", "joy", "laugh", "tear"}, Emoji: "😂", Emoticon: ":')"},
	{Label: "slightly smiling face", Tags: []string{"face", "smile"}, Emoji: "🙂", Emoticon: ":)"},
	{Label: "upside-down face", Tags: []string{"face", "upside-down"}, Emoji: "🙃"},
	{Label: "grapes", Tags: []string{"fruit", "grape"}, Emoji: "🍇"},
	{Label: "melon", Tags: []string{"fruit"}, Emoji: "🍈"},
	{Label: "watermelon", Tags: []string{"fruit"}, Emoji: "🍉"},
	{Label: "tangerine", Tags: []string{"fruit", "orange"}, Emoji: "🍊"},
	{Label: "lemon", Tags: []string{"citrus", "fruit"}, Emoji: "🍋"},
	{Label: "national park", Tags: []string{"park"}, Emoji: "🏞️"},
	{Label: "stadium", Tags: []string{"stadium"}, Emoji: "🏟️"},
	{Label: "locomotive", Tags: []string{"engine", "railway", "steam", "train"}, Emoji: "🚂"},
	{Label: "railway car", Tags: []string{"car", "electric", "railway", "train", "tram", "trolleybus"}, Emoji: "🚃"},
	{Label: "alarm clock", Tags: []string{"alarm", "clock"}, Emoji: "⏰️"},
	{Label: "monkey face", Tags: []string{"face", "monkey"}, Emoji: "🐵"},
	{Label: "dog face", Tags: []string{"dog", "face", "pet"}, Emoji: "🐶"},
	{Label: "fox", Tags: []string{"face"}, Emoji: "🦊"},
	{Label: "cat face", Tags: []string{"cat", "face", "pet"}, Emoji: "🐱"},
	{Label: "tiger face", Tags: []string{"face", "tiger"}, Emoji: "🐯"},
}
