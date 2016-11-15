from django import forms
class VoterForm(forms.ModelForm):
    class Meta:
        model = Voter
        widgets = {
        'passwd': forms.PasswordInput(),
    }
