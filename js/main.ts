const body = document.getElementsByTagName("body")[0]

fetch("/api/photos").then((resp) => {
	return resp.json() as Promise<string[]>
}).then((photos) => {
	for(let p of photos) {
		body.append(createImage(p))
	}
})

function createImage(path: string): HTMLImageElement {
	const img = document.createElement('img')
	img.src = `/images/${path}`
	img.alt = `Photo ${path}`
	return img
}
