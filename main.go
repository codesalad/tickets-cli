package main

import ( "fmt"
	"os"
	"strings"
	"encoding/json"
	"time"
	"flag"
	requests "github.com/codesalad/tickets-cli/requests"
 	"github.com/joho/godotenv"
)

type DatanodeSprints struct {
	Data []Sprint `json:"data"`
}

type DatanodeTickets struct {
	Data []Ticket `json:"data"`
}

type Sprint struct {
	Name string `json:"name"`
	Start string `json:"start"`
	End string `json:"end"`
}

type Ticket struct {
	Id	string `json:"_id,omitempty"`
	Index int `json:"_index"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Assignee string `json:"assignee,omitempty"`
	Status string `json:"sprintStatus,omitempty"`
	TicketStatus string `json:"status,omitempty"`
	Estimated interface{} `json:"estimated,omitempty"`
	Author string `json:"author,omitempty"`
	Sprint string `json:"sprint,omitempty"`
	ChangedBy string `json:"changedBy,omitempty"`
	Archived interface{} `json:"archived,omitempty"`
}

func (t Ticket) String() string {
	assignee := strings.Split(t.Assignee, "@")[0]
	author := strings.Split(t.Author, "@")[0]
	changedBy := strings.Split(t.ChangedBy, "@")[0]
	var people string
	if assignee == author {
		people = ""
	} else if changedBy != "" {
		people = fmt.Sprintf("(%s   %s   %s)", author, changedBy, assignee)
	} else {
		people = fmt.Sprintf("(%s   %s)", author, assignee)
	}

	var archiveMark string
	if t.Archived != nil && t.Archived.(bool) {
		archiveMark = "(archived)"
	} else {
		archiveMark = ""
	}

	const start = "\033]8;;"
    const middle = "\033\\"
    const end = "\033]8;;\033\\"
	var url = os.Getenv("ROOT_URL") +  "/app/Ticket-Service-V2/ticketsingle?ticketId="+t.Id


	link := fmt.Sprintf("%s%s%s%s%s",start,url,middle,t.Name,end)

	return fmt.Sprintf("#%d %s [%s] %s %s", 
		t.Index, link, t.Status, people, archiveMark)
}

func (t Ticket) vString() string {
	assignee := strings.Split(t.Assignee, "@")[0]
	author := strings.Split(t.Author, "@")[0]
	changedBy := strings.Split(t.ChangedBy, "@")[0]
	var people string
	if assignee == author {
		people = ""
	} else if changedBy != "" {
		people = fmt.Sprintf("(%s   %s   %s)", author, changedBy, assignee)
	} else {
		people = fmt.Sprintf("(%s   %s)", author, assignee)
	}
	const start = "\033]8;;"
    const middle = "\033\\"
    const end = "\033]8;;\033\\"
	var url = os.Getenv("ROOT_URL") +  "/app/Ticket-Service-V2/ticketsingle?ticketId="+t.Id

	var archiveMark string
	if t.Archived != nil && t.Archived.(bool) {
		archiveMark = "(archived)"
	} else {
		archiveMark = ""
	}

	link := fmt.Sprintf("%s%s%s%s%s",start,url,middle,t.Name,end)

	return fmt.Sprintf("----------------------\n#%d %s (%vh) [%s] %s %s\n----------------------\n%s\n", 
		t.Index, link, t.Estimated, t.Status, people, archiveMark, t.Description)
}


// func to get sprints
func getSprints() []Sprint {
	var responseData []DatanodeSprints

	ROOT_URL := os.Getenv("ROOT_URL")
	headers := make(map[string]string)
	
	headers["Authorization"] = "ApiKey " + os.Getenv("API_KEY")
	headers["Content-Type"] = "application/json"
	
	r, _ := requests.Get(ROOT_URL + "/data/sprints?pagination=0", headers)
	err := json.Unmarshal([]byte(r), &responseData)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return nil
	}


	return responseData[0].Data
}

func getCurrentSprint() string {
	sprints := getSprints()
	currentTime := time.Now()
	currentDate := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), 0, 0, 0, 0, currentTime.Location())

	for _, sprint := range sprints {
		shortFormat := "2006-01-02"
		startTime, err := time.Parse(shortFormat, sprint.Start)
		if err != nil {
			fmt.Println("Error parsing start time:", err)
			continue
		}

		endTime, err := time.Parse(shortFormat, sprint.End)
		if err != nil {
			fmt.Println("Error parsing end time:", err)
			continue
		}

		if (currentDate.After(startTime) || currentDate.Equal(startTime)) && (currentDate.Before(endTime) || currentDate.Equal(endTime)) {
			return sprint.Name
		}
	}

	return ""
}

func getTickets(sprint string, assignee string, sprintStatus string) []Ticket {
	var responseData []DatanodeTickets

	ROOT_URL := os.Getenv("ROOT_URL")
	headers := make(map[string]string)
	
	headers["Authorization"] = "ApiKey " + os.Getenv("API_KEY")
	headers["Content-Type"] = "application/json"

	if !strings.Contains(assignee, "@serviceheroes.com") {
		assignee = "~" + assignee
	}

	queries := fmt.Sprintf("?archived=false&sprint=%s&assignee=%s&sprintStatus=%s&pagination=0", sprint, assignee, sprintStatus)
	queries = strings.ReplaceAll(queries, " ", "%20")
	
	r, _ := requests.Get(ROOT_URL + "/synergy/data/tickets" + queries, headers)
	err := json.Unmarshal([]byte(r), &responseData)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return nil
	}

	tickets := responseData[0].Data
	
	if len(tickets) == 0 {
		fmt.Println("No tickets found for the given criteria.")
		return nil
	}

	return tickets
}

func printTickets(sprint string, assignee string, status string, verbose bool) {
	tickets := getTickets(sprint, assignee, status)

	bucketedTickets := make(map[string][]Ticket)

	for _, ticket := range tickets {
		bucketedTickets[ticket.Status] = append(bucketedTickets[ticket.Status], ticket)
	}

	for status, tickets := range bucketedTickets {
		fmt.Printf("\n============= %s =============\n", status)
		for _, ticket := range tickets {
			if verbose {
				fmt.Println(ticket.vString())
				continue
			}
			fmt.Println(ticket)
		}
	}


}

func createTicket(user string, sprint string, assignee string, status string, name string, description string) {
	newTicket := Ticket{
		Name: name,
		Description: description,
		Assignee: assignee,
		Status: status,
		TicketStatus: "Backlog",
		Sprint: sprint,
		Author: user,
		ChangedBy: user,
		Archived: "false",
	}

	fmt.Println("Creating ticket:", newTicket)

	ROOT_URL := os.Getenv("ROOT_URL")
	headers := make(map[string]string)
	
	headers["Authorization"] = "ApiKey " + os.Getenv("API_KEY")
	headers["Content-Type"] = "application/json"

	r, _ := requests.Post(ROOT_URL + "/synergy/data/tickets", newTicket, headers)
	fmt.Println("Response Body:", r)
}

func updateTicket(index int, ticket Ticket) {
	ROOT_URL := os.Getenv("ROOT_URL")
	headers := make(map[string]string)
	
	headers["Authorization"] = "ApiKey " + os.Getenv("API_KEY")
	headers["Content-Type"] = "application/json"
	
	r, s := requests.Patch(fmt.Sprintf("%s/synergy/data/tickets?_index=%d", ROOT_URL, index), ticket, headers)
	fmt.Println("Response Status:", s)
	fmt.Println("Response Body:", r)
}

func deleteTicket(index int) {
	ROOT_URL := os.Getenv("ROOT_URL")
	headers := make(map[string]string)
	
	headers["Authorization"] = "ApiKey " + os.Getenv("API_KEY")
	headers["Content-Type"] = "application/json"
	
	r, s := requests.Delete(fmt.Sprintf("%s/synergy/data/tickets?_index=%d", ROOT_URL, index), headers)
	fmt.Println("Response Status:", s)
	fmt.Println("Response Body:", r)
}


func main() {
	err := godotenv.Load(fmt.Sprintf("%s/.config/tickets.config", os.Getenv("HOME")))
	if err != nil {
		fmt.Println("Error loading .env file. Make ure tickets.config exists in ~/.config")
		return
	}

	currentSprint := flag.String("sprint", getCurrentSprint(), "Current Sprint Name" )
	user := flag.String("username", os.Getenv("TICKETS_USERNAME"), "Username")
	create := flag.Bool("create", false, "Create a ticket")

	del := flag.Int("delete", -1, "Delete a ticket by its index")
	update := flag.Int("update", -1, "Update a ticket by its index")
	assignee := flag.String("assignee", os.Getenv("TICKETS_USERNAME"), "Assignee to view or create ticket for")
	status := flag.String("status", "", "Sprint status to filter tickets (Backlog, Next Week, To do, In Progress, Done)")

	verbose := flag.Bool("v", false, "Enable verbose output")
	
	name := flag.String("name", "", "Name of the ticket to create")
	description := flag.String("description", "", "Description of the ticket to create")

	// -- short handles
	flag.BoolVar(create, "c", false, "Create a ticket")
	flag.IntVar(del, "d", -1, "Delete a ticket by its index")
	flag.IntVar(update, "u", -1, "Update a ticket by its index")

	flag.Parse()
	fmt.Printf("Selected Sprint: %s\n", *currentSprint)
	fmt.Printf("User: %s\n", *user)
	if *create {
		if *del != -1 {
			fmt.Println("Error: --create and --delete cannot be used together.")
			return
		}
		if *update != -1 {
			fmt.Println("Error: --create and --update cannot be used together.")
			return
		}
		if *status == "" {
			*status = "To do"
		}
		if *name == "" {
			fmt.Println("Error: --name is required to create a ticket.")
			return
		}
		createTicket(*user, *currentSprint, *assignee, *status, *name, *description)
	} else if *del != -1 {
		if *update != -1 {
			fmt.Println("Error: --delete and --update cannot be used together.")
			return
		}
		//deleteTicket(*del) // Disabled for safety
		patchTicket := Ticket{
			Archived: "true",
		}
		updateTicket(*del, patchTicket)

	} else if *update != -1 {
		patchTicket := Ticket{}
		if *name != "" {
			patchTicket.Name = *name
		}
		if *description != "" {
			patchTicket.Description = *description
		}
		if *assignee != "" {
			patchTicket.Assignee = *assignee
		}
		if *status != "" {
			patchTicket.Status = *status
		}
		if *currentSprint != "" {
			patchTicket.Sprint = *currentSprint
		}
		patchTicket.ChangedBy = *user

		updateTicket(*update, patchTicket)
	} else {
		if *status == "" {	
			*status = "To do;In Progress"
		}
		printTickets(*currentSprint, *assignee, *status, *verbose)
	}

}
