import pytest
import respx
from httpx import Response
from ticket_client.client import TicketClient
from ticket_client.models import TicketInput

@respx.mock
def test_handle_null_and_empty_area():
    client = TicketClient()
    
    # Mock response with null and empty area
    respx.get("http://localhost:8080/tickets").mock(return_value=Response(200, json=[
        {"id": 1, "title": "Null Area", "area_of_concern": None, "status": "Open"},
        {"id": 2, "title": "Empty Area", "area_of_concern": "", "status": "Open"}
    ]))
    
    tickets = client.list_tickets()
    assert len(tickets) == 2
    assert tickets[0].area_of_concern is None
    assert tickets[1].area_of_concern == ""

@respx.mock
def test_empty_string_filter_passed_to_server():
    client = TicketClient()
    
    # Verify that exclude_area="" results in a query param
    route = respx.get("http://localhost:8080/tickets").mock(return_value=Response(200, json=[]))
    
    client.list_tickets(exclude_area="")
    
    assert route.called
    assert route.calls.last.request.url.params["exclude_area"] == ""

@respx.mock
def test_long_resolution_message_fidelity():
    client = TicketClient()
    long_msg = "Resolved" * 1000
    
    respx.get("http://localhost:8080/tickets/1").mock(return_value=Response(200, json={
        "id": 1, "title": "Long", "resolution_message": long_msg, "status": "Closed"
    }))
    
    ticket = client.get_ticket(1)
    assert ticket.resolution_message == long_msg

@respx.mock
def test_empty_resolution_message():
    client = TicketClient()
    
    respx.get("http://localhost:8080/tickets/1").mock(return_value=Response(200, json={
        "id": 1, "title": "Empty Res", "resolution_message": "", "status": "Closed"
    }))
    
    ticket = client.get_ticket(1)
    assert ticket.resolution_message == ""
