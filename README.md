# tickets-cli

## Setup
1. Ensure you have Go installed on your machine. You can download it from [https://golang.org/dl/](https://golang.org/dl/).
2. Install the `tickets-cli` tool using the following command:
    ```
    go install github.com/codesalad/tickets-cli@latest
    ```
3. Make sure your Go bin directory is in your system's PATH. You can add the following line to your shell configuration file (e.g., `.bashrc`, `.zshrc`):
    ```
    export PATH=$PATH:$(go env GOPATH)/bin
    ```
4. Restart your terminal or run `source ~/.bashrc` (or the appropriate file for your shell) to apply the changes.   
5. Initialize the config file in $HOME/.config/tickets.config
    ```
    ROOT_URL="root url of the ticket tool"
    API_KEY="your api key with data scope"
    USERNAME="your email address"
    ```

    

## Example Commands
- View all tickets assigned to you:
    ```
    $: tickets
    ```

- View tickets assigned to a specific user:
    ```
    $: tickets -assignee "user@domain.com"
    ```

- Create a new ticket:
    ```
    $: tickets -c -name "New Feature" -description "Description of the new feature"
    ```
- Update an existing ticket:
    ```
    $: tickets -u 1 -name "Updated Feature Name" -description "Updated description"
    ```
- Delete a ticket:
    ```
    $: tickets -d 1
    ```
- View tickets in a specific sprint:
    ```
    $: tickets -sprint "2025-10"
    ```
- View tickets with a specific status:`
    ```
    $: tickets -status "In Progress"
    ```
- Enable verbose output for debugging:
    ```
    $: tickets -v
    ```

## Usage
  -u int
    	Update a ticket by its index (default -1)
  -assignee string
    	Assignee to view or create ticket for (default "wing@serviceheroes.com")
  -c	Create a ticket
  -create
    	Create a ticket
  -d int
    	Delete a ticket by its index (default -1)
  -delete int
    	Delete a ticket by its index (default -1)
  -description string
    	Description of the ticket to create
  -name string
    	Name of the ticket to create
  -sprint string
    	Current Sprint Name (default "2025-10")
  -status string
    	Sprint status to filter tickets (Backlog, Next Week, To do, In Progress, Done)
  -update int
    	Update a ticket by its index (default -1)
  -username string
    	Username (default "wing@serviceheroes.com")
  -v	Enable verbose output

