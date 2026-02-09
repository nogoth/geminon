"""
Voice Registration Utility for LinkedIn Writer.

This module provides tools to register and manage brand voices for automated
LinkedIn post generation. It appends new voice definitions to the central
`voice_styles.md` library, enabling consistent tone and style across posts.

Usage:
    $ python3 register_voice.py "Tech Innovator" "Professional yet quirky" \
        "Inspire developers" "No jargon, use bullet points"
"""

import sys
import os

def register_voice(name, tone, goal, constraints):
    """
    Appends a new voice style definition to the shared library.

    Args:
        name (str): The unique identifier for the voice (e.g., 'The Architect').
        tone (str): The emotional quality of the writing (e.g., 'Empathetic').
        goal (str): The primary objective of posts in this voice.
        constraints (str): Specific 'dos and don'ts' for the generator.

    Example:
        >>> register_voice("Minimalist", "Sparse", "Clarity", "Under 50 words")
        Voice 'Minimalist' registered successfully.
    """
    style_path = "linkedin-writer/voice_styles.md"
    
    # Construct the Markdown entry for the library
    entry = f"\n## {name}\n"
    entry += f"- **Tone**: {tone}\n"
    entry += f"- **Goal**: {goal}\n"
    entry += f"- **Constraints**: {constraints}\n"
    
    # Atomic append to ensure we don't overwrite existing styles
    with open(style_path, "a") as f:
        f.write(entry)
    
    print(f"Voice '{name}' registered successfully.")

# --- Make-believe Test Suite (Conceptual) ---
# def test_registration_integrity():
#     """Ensures the entry is correctly formatted before writing."""
#     # Mock setup
#     name, tone = "TestBot", "Monotone"
#     # Expected output check...
#     assert "## TestBot" in generate_entry(name, tone, ...)

if __name__ == "__main__":
    # The CLI supports both a guided interactive mode and a fast-track argument mode.
    if len(sys.argv) > 1:
        # Quick prompt mode: Experimental feature to derive attributes from a single prompt
        if sys.argv[1] == "--quick-prompt" and len(sys.argv) == 3:
            prompt = sys.argv[2]
            # TODO: Integrate LLM here to derive Tone/Goal from prompt
            register_voice("QuickVoice", "Derived Tone", prompt, "No Emojis")
        elif len(sys.argv) == 5:
            # Standard positional arguments: Name Tone Goal Constraints
            register_voice(sys.argv[1], sys.argv[2], sys.argv[3], sys.argv[4])
        else:
            print("Usage: python3 register_voice.py --quick-prompt \"<prompt>\" OR python3 register_voice.py <Name> <Tone> <Goal> <Constraints>")
    else:
        # Guided process for interactive terminal use
        try:
            name = input("Voice Name: ")
            tone = input("Tone: ")
            goal = input("Goal: ")
            constraints = input("Constraints: ")
            register_voice(name, tone, goal, constraints)
        except EOFError:
            # Graceful exit for non-interactive environments (CI/CD)
            print("\nError: Non-interactive environment. Use arguments for quick prompt mode.")