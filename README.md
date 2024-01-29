# Desc
Using bubbletea and libgloss to write a web favorates handler for command line usage.

Multiple ways to sort: label, folder

Concept image:

# Targets

    Have a nice UI, customization of UI components'position, maybe drag with mouse and store the new position parameter
    Add/delete web
    Description for each web
    Multiple name alias for each web page
    Scroll bar
    Open a web page through target web browser
    Icon setting for each web page
    Fuzzy finder
    Classification support
    Vim mode for moving around and search space, maybe customization for keybindings
    YMAL configuration file
    Custom sort way: folder, label
    Flip between description and link
    Copy the link to clipboard
    Edit link and description in real time
    Basic help info on the top, could be hidden
    Check whether the keybindings is conflict
    Configuration file should be placed at "~/.config/web-tree/"
    Every tree have a storage file? data 目录, 每创建一个 tree, 就添加一个文件
    Check configuration

Possible TUI for it:
- text input
- help
- key
- autocomplete
- composable-views
- http
- list-fancy
- progress-animated
- spinners
- table
- tabs
- tui-daemon-combo

1. Command line args
2. Draw the outer box
3. help, key menu
4. text input and autocomplete for search space
5. Add function for http testing
# Usage
## Move around

    `?`: Help manual
    `<Tab>`: Moving between several space
    `j`: Move down
    `k`: Move up
    `h`: Move left
    `l`: Move right
    `o`: Open in web browser
    `a`: Add new web
    `d`: Delete a selected web
    `t`: Toggle from description and web link

## Work space

    Search, on the top of the UI
    Web tree, 




