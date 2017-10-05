package blobstoreexample

import (
	"encoding/json"
	"io"
	"net/http"

	"golang.org/x/net/context"
	"google.golang.org/appengine"
	"google.golang.org/appengine/blobstore"
	"google.golang.org/appengine/log"
)

//ResponseTemp Temporário
type ResponseTemp struct {
	Host string
	Path string
}

func serveError(ctx context.Context, w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Header().Set("Content-Type", "text/plain")
	io.WriteString(w, "Internal Server Error")
	log.Errorf(ctx, "%v", err)
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	uploadURL, err := blobstore.UploadURL(ctx, "/upload", nil)
	if err != nil {
		serveError(ctx, w, err)
		return
	}
	w.Header().Set("Content-Type", "application/json")

	log.Infof(ctx, "Url Path %s", uploadURL.Path)

	responsePartial := &ResponseTemp{Path: uploadURL.Path, Host: uploadURL.Host}

	js, err := json.Marshal(responsePartial)

	log.Infof(ctx, "Objeto Response %s", js)

	w.Write(js)
	if err != nil {
		log.Errorf(ctx, "%v", err)
	}
}

func handleServe(w http.ResponseWriter, r *http.Request) {
	blobstore.Send(w, appengine.BlobKey(r.FormValue("blobKey")))
}

func handleUpload(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	blobs, _, err := blobstore.ParseUpload(r)
	if err != nil {
		serveError(ctx, w, err)
		return
	}
	file := blobs["file"]
	if len(file) == 0 {
		log.Errorf(ctx, "no file uploaded")
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}
	//	http.Redirect(w, r, "/serve/?blobKey="+string(file[0].BlobKey), http.StatusFound)
}

func init() {
	http.HandleFunc("/", handleRoot)
	http.HandleFunc("/serve/", handleServe)
	http.HandleFunc("/upload", handleUpload)
}
