import typer
from typing import Optional
from rich.console import Console
from rich.table import Table
from rich.panel import Panel
from rich import print as rprint
from .client import TicketClient
from .models import TicketInput

app = typer.Typer(help="Ticket Tracker CLI Client")
console = Console()
client = TicketClient()

@app.command()
def create(
    title: str,
    description: str = "",
    area: str = "General",
    agent: str = "Unassigned",
    status: str = "Open"
):
    """Create a new ticket."""
    ticket_in = TicketInput(
        title=title,
        description=description,
        area_of_concern=area,
        agent_name=agent,
        status=status
    )
    try:
        ticket = client.create_ticket(ticket_in)
        rprint(Panel(f"Ticket Created Successfully! ID: [bold green]{ticket.id}[/]", title="Success"))
    except Exception as e:
        rprint(f"[bold red]Error:[/] {e}")

@app.command()
def list(
    exclude_area: Optional[str] = typer.Option(None, "--exclude-area", help="Filter out area"),
    exclude_status: Optional[str] = typer.Option(None, "--exclude-status", help="Filter out status")
):
    """List all tickets with optional filters."""
    try:
        tickets = client.list_tickets(exclude_area=exclude_area, exclude_status=exclude_status)
        if not tickets:
            rprint("[yellow]No tickets found.[/]")
            return

        table = Table(title="Tickets")
        table.add_column("ID", style="cyan")
        table.add_column("Title", style="magenta")
        table.add_column("Status", style="green")
        table.add_column("Area", style="blue")
        table.add_column("Agent", style="yellow")

        for t in tickets:
            table.add_row(str(t.id), t.title, t.status, t.area_of_concern, t.agent_name)
        
        console.print(table)
    except Exception as e:
        rprint(f"[bold red]Error:[/] {e}")

@app.command()
def get(ticket_id: int):
    """Get details of a specific ticket."""
    try:
        t = client.get_ticket(ticket_id)
        content = f"[bold]Title:[/] {t.title}\n"
        content += f"[bold]Status:[/] {t.status}\n"
        content += f"[bold]Area:[/] {t.area_of_concern}\n"
        content += f"[bold]Agent:[/] {t.agent_name}\n"
        content += f"[bold]Description:[/] {t.description}\n"
        content += f"[bold]Resolution:[/] {t.resolution_message or 'N/A'}\n"
        content += f"[bold]Created:[/] {t.created_at}\n"
        content += f"[bold]Updated:[/] {t.updated_at}"
        
        rprint(Panel(content, title=f"Ticket #{t.id}", expand=False))
    except Exception as e:
        rprint(f"[bold red]Error:[/] {e}")

@app.command()
def resolve(ticket_id: int, message: str):
    """Resolve a ticket with a message."""
    try:
        # First get existing
        t = client.get_ticket(ticket_id)
        # Update
        ticket_in = TicketInput(
            title=t.title,
            description=t.description,
            area_of_concern=t.area_of_concern,
            agent_name=t.agent_name,
            status="Closed",
            resolution_message=message
        )
        updated = client.update_ticket(ticket_id, ticket_in)
        rprint(f"[bold green]Ticket #{updated.id} resolved![/]")
    except Exception as e:
        rprint(f"[bold red]Error:[/] {e}")

if __name__ == "__main__":
    app()
