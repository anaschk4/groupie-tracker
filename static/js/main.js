function performSearch() {
    const query = document.getElementById('search-bar').value;
    const creationDate = document.getElementById('creation-date').value;
    const firstAlbum = document.getElementById('first-album').value;

    const params = new URLSearchParams({
        query: query,
        creationDate: creationDate,
        firstAlbum: firstAlbum
    });

    console.log("URL générée :", `/search?${params.toString()}`);

    fetch(`/search?${params.toString()}`, {
        headers: { "X-Requested-With": "XMLHttpRequest" }
    })
        .then(response => {
            if (!response.ok) {
                throw new Error("Erreur réseau : " + response.statusText);
            }
            return response.json();
        })
        .then(data => {
            const resultsDiv = document.getElementById('results-container');
            resultsDiv.innerHTML = '';

            if (data.length === 0) {
                resultsDiv.innerHTML = '<p>Aucun résultat trouvé.</p>';
                return;
            }

            data.forEach(item => {
                const div = document.createElement('div');
                div.className = 'artist-card';
                div.innerHTML = `
                    <img src="${item.Image}" alt="${item.Name}" class="artist-image">
                    <div class="artist-info">
                        <h2 class="artist-name">${item.Name}</h2>
                        <div class="basic-info">
                            <p><strong>Date de création:</strong> ${item.CreationDate}</p>
                            <p><strong>Premier album:</strong> ${item.FirstAlbum}</p>
                        </div>
                        <div class="info-section">
                            <h3>Membres</h3>
                            <ul class="info-list">
                                ${item.Members.map(member => `<li>${member}</li>`).join('')}
                            </ul>
                        </div>
                    </div>
                `;
                resultsDiv.appendChild(div);
            });
        })
        .catch(error => {
            console.error("Erreur détectée :", error);
            const resultsDiv = document.getElementById('results-container');
            resultsDiv.innerHTML = '<p>Une erreur est survenue. Veuillez réessayer.</p>';
        });
}
