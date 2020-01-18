# Zargon

Zargon is a utility to set-up and manage your story.

## Usage

### `zargon init [options] theme`

Initializes a directory to be used as a Zargon story.

Based on your configuration is will download a number of random generators for `npcs`, `quests`, `items` and `places` based on the given theme.
This means that each time your run `zargon init` you'll set an entirely different adventure!

Note that Zargon GM works with any plain text file, anywhere on your disk, in any file structure.
The `zargon init` command is strictly a convenience and not in any way a neccesity.

```
$ zargin init --help
usage: zargon init [options] theme
    options:
        -s, --seed  seed used to randomize generators
$ zargin init fantasy
$ tree
.
├── gen
│   ├── npc.gen
│   ├── quest.gen
│   ├── item.gen
│   └── place.gen
```

### `zargon gen`

Manages your installed generators.

#### `zargon gen list`

Lists all available generators on the default generator marketplace.

```
$ zargon gen list --help
usage: zargon gen list
$ zargon gen list
fantasy_theme_npc_plain.gen@v0.0.1 {fantasy}       [NPC]   Generates plain medieval fantasy NPC's.
fantasy_theme_npc_dark.gen@v0.0.1  {fantasy, dark} [NPC]   Generates dark (evil, mean spirited, mysterious) medieval fantasy NPC's.
scifi_theme_npc_plain.gen@v0.0.1   {scifi}         [NPC]   Generates plain sci-fi fantasy NPC's.
horror_theme_npc_plain.gen@v0.0.1  {horror, dark}  [NPC]   Generates villainous horror NPC's.
```

#### `zargon gen install`

Installs an available generator on the default generator marketplace.

```
$ zargon gen install --help
usage: zargon gen install [options] generator_name
    options:
        -o, --output_dir    directory to download the generator to, defaults to `./gen`
```

#### `zargon gen uninstall`

Uninstalls an installed generator.

```
$ zargon gen uninstall --help
usage: zargon gen uninstall generator_name
```
