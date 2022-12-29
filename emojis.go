package main

import (
	"fmt"

	"github.com/krish-r/go-commit/color"
)

type Emoji struct {
	id          int
	name        string
	emoji       string
	code        string
	description string
}

func (e Emoji) print() {
	id := color.Colorize(e.id, color.BrightGreen)
	description := color.Colorize("("+e.description+")", color.BrightBlackGray)

	fmt.Printf("%s: %s %s %s\n", id, e.name, e.emoji, description)
}

var emojis = []Emoji{
	{id: 0, name: "none", emoji: "", code: "", description: "none"},
	{id: 1, name: "art", emoji: "🎨", code: ":art:", description: "Improve structure / format of the code."},
	{id: 2, name: "zap", emoji: "⚡️", code: ":zap:", description: "Improve performance."},
	{id: 3, name: "fire", emoji: "🔥", code: ":fire:", description: "Remove code or files."},
	{id: 4, name: "bug", emoji: "🐛", code: ":bug:", description: "Fix a bug."},
	{id: 5, name: "ambulance", emoji: "🚑️", code: ":ambulance:", description: "Critical hotfix."},
	{id: 6, name: "sparkles", emoji: "✨", code: ":sparkles:", description: "Introduce new features."},
	{id: 7, name: "memo", emoji: "📝", code: ":memo:", description: "Add or update documentation."},
	{id: 8, name: "rocket", emoji: "🚀", code: ":rocket:", description: "Deploy stuff."},
	{id: 9, name: "lipstick", emoji: "💄", code: ":lipstick:", description: "Add or update the UI and style files."},
	{id: 10, name: "tada", emoji: "🎉", code: ":tada:", description: "Begin a project."},
	{id: 11, name: "white-check-mark", emoji: "✅", code: ":white_check_mark:", description: "Add, update, or pass tests."},
	{id: 12, name: "lock", emoji: "🔒️", code: ":lock:", description: "Fix security issues."},
	{id: 13, name: "closed-lock-with-key", emoji: "🔐", code: ":closed_lock_with_key:", description: "Add or update secrets."},
	{id: 14, name: "bookmark", emoji: "🔖", code: ":bookmark:", description: "Release / Version tags."},
	{id: 15, name: "rotating-light", emoji: "🚨", code: ":rotating_light:", description: "Fix compiler / linter warnings."},
	{id: 16, name: "construction", emoji: "🚧", code: ":construction:", description: "Work in progress."},
	{id: 17, name: "green-heart", emoji: "💚", code: ":green_heart:", description: "Fix CI Build."},
	{id: 18, name: "arrow-down", emoji: "⬇️", code: ":arrow_down:", description: "Downgrade dependencies."},
	{id: 19, name: "arrow-up", emoji: "⬆️", code: ":arrow_up:", description: "Upgrade dependencies."},
	{id: 20, name: "pushpin", emoji: "📌", code: ":pushpin:", description: "Pin dependencies to specific versions."},
	{id: 21, name: "construction-worker", emoji: "👷", code: ":construction_worker:", description: "Add or update CI build system."},
	{id: 22, name: "chart-with-upwards-trend", emoji: "📈", code: ":chart_with_upwards_trend:", description: "Add or update analytics or track code."},
	{id: 23, name: "recycle", emoji: "♻️", code: ":recycle:", description: "Refactor code."},
	{id: 24, name: "heavy-plus-sign", emoji: "➕", code: ":heavy_plus_sign:", description: "Add a dependency."},
	{id: 25, name: "heavy-minus-sign", emoji: "➖", code: ":heavy_minus_sign:", description: "Remove a dependency."},
	{id: 26, name: "wrench", emoji: "🔧", code: ":wrench:", description: "Add or update configuration files."},
	{id: 27, name: "hammer", emoji: "🔨", code: ":hammer:", description: "Add or update development scripts."},
	{id: 28, name: "globe-with-meridians", emoji: "🌐", code: ":globe_with_meridians:", description: "Internationalization and localization."},
	{id: 29, name: "pencil2", emoji: "✏️", code: ":pencil2:", description: "Fix typos."},
	{id: 30, name: "poop", emoji: "💩", code: ":poop:", description: "Write bad code that needs to be improved."},
	{id: 31, name: "rewind", emoji: "⏪️", code: ":rewind:", description: "Revert changes."},
	{id: 32, name: "twisted-rightwards-arrows", emoji: "🔀", code: ":twisted_rightwards_arrows:", description: "Merge branches."},
	{id: 33, name: "package", emoji: "📦️", code: ":package:", description: "Add or update compiled files or packages."},
	{id: 34, name: "alien", emoji: "👽️", code: ":alien:", description: "Update code due to external API changes."},
	{id: 35, name: "truck", emoji: "🚚", code: ":truck:", description: "Move or rename resources (e.g.: files, paths, routes)."},
	{id: 36, name: "page-facing-up", emoji: "📄", code: ":page_facing_up:", description: "Add or update license."},
	{id: 37, name: "boom", emoji: "💥", code: ":boom:", description: "Introduce breaking changes."},
	{id: 38, name: "bento", emoji: "🍱", code: ":bento:", description: "Add or update assets."},
	{id: 39, name: "wheelchair", emoji: "♿️", code: ":wheelchair:", description: "Improve accessibility."},
	{id: 40, name: "bulb", emoji: "💡", code: ":bulb:", description: "Add or update comments in source code."},
	{id: 41, name: "beers", emoji: "🍻", code: ":beers:", description: "Write code drunkenly."},
	{id: 42, name: "speech-balloon", emoji: "💬", code: ":speech_balloon:", description: "Add or update text and literals."},
	{id: 43, name: "card-file-box", emoji: "🗃️", code: ":card_file_box:", description: "Perform database related changes."},
	{id: 44, name: "loud-sound", emoji: "🔊", code: ":loud_sound:", description: "Add or update logs."},
	{id: 45, name: "mute", emoji: "🔇", code: ":mute:", description: "Remove logs."},
	{id: 46, name: "busts-in-silhouette", emoji: "👥", code: ":busts_in_silhouette:", description: "Add or update contributor(s)."},
	{id: 47, name: "children-crossing", emoji: "🚸", code: ":children_crossing:", description: "Improve user experience / usability."},
	{id: 48, name: "building-construction", emoji: "🏗️", code: ":building_construction:", description: "Make architectural changes."},
	{id: 49, name: "iphone", emoji: "📱", code: ":iphone:", description: "Work on responsive design."},
	{id: 50, name: "clown-face", emoji: "🤡", code: ":clown_face:", description: "Mock things."},
	{id: 51, name: "egg", emoji: "🥚", code: ":egg:", description: "Add or update an easter egg."},
	{id: 52, name: "see-no-evil", emoji: "🙈", code: ":see_no_evil:", description: "Add or update a .gitignore file."},
	{id: 53, name: "camera-flash", emoji: "📸", code: ":camera_flash:", description: "Add or update snapshots."},
	{id: 54, name: "alembic", emoji: "⚗️", code: ":alembic:", description: "Perform experiments."},
	{id: 55, name: "mag", emoji: "🔍️", code: ":mag:", description: "Improve SEO."},
	{id: 56, name: "label", emoji: "🏷️", code: ":label:", description: "Add or update types."},
	{id: 57, name: "seedling", emoji: "🌱", code: ":seedling:", description: "Add or update seed files."},
	{id: 58, name: "triangular-flag-on-post", emoji: "🚩", code: ":triangular_flag_on_post:", description: "Add, update, or remove feature flags."},
	{id: 59, name: "goal-net", emoji: "🥅", code: ":goal_net:", description: "Catch errors."},
	{id: 60, name: "animation", emoji: "💫", code: ":dizzy:", description: "Add or update animations and transitions."},
	{id: 61, name: "wastebasket", emoji: "🗑️", code: ":wastebasket:", description: "Deprecate code that needs to be cleaned up."},
	{id: 62, name: "passport-control", emoji: "🛂", code: ":passport_control:", description: "Work on code related to authorization, roles and permissions."},
	{id: 63, name: "adhesive-bandage", emoji: "🩹", code: ":adhesive_bandage:", description: "Simple fix for a non-critical issue."},
	{id: 64, name: "monocle-face", emoji: "🧐", code: ":monocle_face:", description: "Data exploration/inspection."},
	{id: 65, name: "coffin", emoji: "⚰️", code: ":coffin:", description: "Remove dead code."},
	{id: 66, name: "test-tube", emoji: "🧪", code: ":test_tube:", description: "Add a failing test."},
	{id: 67, name: "necktie", emoji: "👔", code: ":necktie:", description: "Add or update business logic"},
	{id: 68, name: "stethoscope", emoji: "🩺", code: ":stethoscope:", description: "Add or update healthcheck."},
	{id: 69, name: "bricks", emoji: "🧱", code: ":bricks:", description: "Infrastructure related changes."},
	{id: 70, name: "technologist", emoji: "🧑‍💻", code: ":technologist:", description: "Improve developer experience"},
	{id: 71, name: "money-with-wings", emoji: "💸", code: ":money_with_wings:", description: "Add sponsorships or money related infrastructure."},
}
