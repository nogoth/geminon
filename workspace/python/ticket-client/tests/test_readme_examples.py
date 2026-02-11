import re
import pytest
import respx
from httpx import Response
from typer.testing import CliRunner
from ticket_client.main import app

runner = CliRunner()

def extract_commands():
    with open("README.md", "r") as f:
        content = f.read()
    
    # Find all 'uv run ticket-client ...' commands in code blocks
    # We look for lines starting with uv run ticket-client
    commands = re.findall(r"uv run ticket-client (.*)", content)
    return [cmd.strip() for cmd in commands]

@respx.mock
def test_readme_commands():
    commands = extract_commands()
    assert len(commands) > 0, "No commands found in README.md"

    # Mock endpoints for README examples
    # 1. Create
    respx.post("http://localhost:8080/tickets").mock(return_value=Response(201, json={
        "id": 1, "title": "Issue Title", "description": "Details", "status": "Open", 
        "area_of_concern": "IT", "agent_name": "Unassigned", "resolution_message": None,
        "created_at": "2026-02-11T00:00:00Z", "updated_at": "2026-02-11T00:00:00Z"
    }))
    
    # 2. List
    respx.get("http://localhost:8080/tickets").mock(return_value=Response(200, json=[
        {"id": 1, "title": "T1", "status": "Open", "area_of_concern": "IT", "agent_name": "A1"}
    ]))
    
    # 3. Get
    respx.get("http://localhost:8080/tickets/1").mock(return_value=Response(200, json={
        "id": 1, "title": "Issue Title", "description": "Details", "status": "Open", 
        "area_of_concern": "IT", "agent_name": "Unassigned", "resolution_message": None,
        "created_at": "2026-02-11T00:00:00Z", "updated_at": "2026-02-11T00:00:00Z"
    }))
    
    # 4. Resolve (PUT)
    respx.put("http://localhost:8080/tickets/1").mock(return_value=Response(200, json={
        "id": 1, "title": "Issue Title", "description": "Details", "status": "Closed", 
        "area_of_concern": "IT", "agent_name": "Unassigned", "resolution_message": "Fixed the issue",
        "created_at": "2026-02-11T00:00:00Z", "updated_at": "2026-02-11T00:00:00Z"
    }))

    for cmd in commands:
        # Split command into parts, handling quotes correctly
        import shlex
        args = shlex.split(cmd)
        
        result = runner.invoke(app, args)
        assert result.exit_code == 0, f"Command failed: ticket-client {cmd}
Output: {result.stdout}"
        print(f"Verified command: ticket-client {cmd}")
