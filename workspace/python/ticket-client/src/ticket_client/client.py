import os
import httpx
from typing import List, Optional
from .models import Ticket, TicketInput

class TicketClient:
    def __init__(self, base_url: Optional[str] = None):
        if base_url is None:
            host = os.getenv("TICKET_SERVER_HOST", "localhost")
            base_url = f"http://{host}:8080"
        self.base_url = base_url.rstrip("/")

    def create_ticket(self, ticket: TicketInput) -> Ticket:
        with httpx.Client() as client:
            response = client.post(
                f"{self.base_url}/tickets",
                json=ticket.model_dump(exclude_none=True)
            )
            response.raise_for_status()
            return Ticket.model_validate(response.json())

    def list_tickets(self, exclude_area: Optional[str] = None, exclude_status: Optional[str] = None) -> List[Ticket]:
        params = {}
        if exclude_area is not None:
            params["exclude_area"] = exclude_area
        if exclude_status is not None:
            params["exclude_status"] = exclude_status

        with httpx.Client() as client:
            response = client.get(f"{self.base_url}/tickets", params=params)
            response.raise_for_status()
            return [Ticket.model_validate(t) for t in response.json()]

    def get_ticket(self, ticket_id: int) -> Ticket:
        with httpx.Client() as client:
            response = client.get(f"{self.base_url}/tickets/{ticket_id}")
            response.raise_for_status()
            return Ticket.model_validate(response.json())

    def update_ticket(self, ticket_id: int, ticket: TicketInput) -> Ticket:
        with httpx.Client() as client:
            response = client.put(
                f"{self.base_url}/tickets/{ticket_id}",
                json=ticket.model_dump(exclude_none=True)
            )
            response.raise_for_status()
            return Ticket.model_validate(response.json())
