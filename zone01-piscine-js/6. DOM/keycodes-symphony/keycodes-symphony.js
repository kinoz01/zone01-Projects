export function compose() {
    document.addEventListener('keydown', (event) => {
        if (event.key >= 'a' && event.key <= 'z') {
            createNote(event.key);
        }
        if (event.key === 'Backspace') {
            removeLastNote();
        }
        if (event.key === 'Escape') {
            clearAllNotes();
        }
    });
}

function createNote(char) {
    const note = document.createElement('div')
    note.classList.add('note')
    note.style.backgroundColor = `rgb(10, ${char.charCodeAt(0)}, 10)`
    note.textContent = char
    document.body.append(note)
}

function removeLastNote() {
    const notes = document.querySelectorAll('.note');
    if (notes.length > 0) {
        notes[notes.length - 1].remove()
    }
}

function clearAllNotes() {
    const notes = document.querySelectorAll('.note')
    notes.forEach(note=> {
        note.classList = ''
    })
}
