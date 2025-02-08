import re

# Basic pattern matching
text = "Python 3.9 was released in 2020"
pattern = r"\d+\.\d+"  # Matches version number

match = re.search(pattern, text)
if match:
    print(f"Found version: {match.group()}")

# Finding all matches
text = "Contact us at info@example.com or support@example.com"
emails = re.findall(r'[\w\.-]+@[\w\.-]+', text)
print(f"\nFound emails: {emails}")

# Substitution
text = "My phone number is 123-456-7890"
censored = re.sub(r'\d{3}-\d{3}-\d{4}', 'XXX-XXX-XXXX', text)
print(f"\nCensored: {censored}")

# Validation
def is_valid_email(email):
    pattern = r'^[\w\.-]+@[\w\.-]+\.\w+$'
    return bool(re.match(pattern, email))

print("\nEmail validation:")
emails = ["user@example.com", "invalid.email@", "another@example"]
for email in emails:
    print(f"{email}: {'Valid' if is_valid_email(email) else 'Invalid'}") 