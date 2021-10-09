package requestHandler

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"path"
    "regexp"
	"os"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"insta_backend/models"
	"insta_backend/encrypt"
)

type userHandler struct{
	col mongo.Collection
	ctx context.Context
}
type postHandler struct{
	col mongo.Collection
	ctx context.Context
}


var (
	getUserRe   =  regexp.MustCompile(`^\/users\/(\d+)$`)
	getUserPostRe    = regexp.MustCompile(`^\/posts\/users\/(\d+)$`)
	getPostRe =  regexp.MustCompile(`^\/posts\/(\d+)$`)
	createUserRe = regexp.MustCompile(`^\/users[\/]*$`)
	createPostRe = regexp.MustCompile(`^\/posts[\/]*$`)
)


func (h *userHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	fmt.Println(r)
	switch{
	case r.Method == http.MethodPost && createUserRe.MatchString(r.URL.Path):
		{
			h.createUser(w, r)
			return
		}
	case r.Method == http.MethodGet && getUserRe.MatchString(r.URL.Path):
		{
			h.getUser(w, r)
			return
		}
	default:
        notFound(w, r)
        return
    }
}



func (h *postHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	switch{
	case r.Method == http.MethodPost && createPostRe.MatchString(r.URL.Path):
		{
			h.createPost(w,r)
			return
		}
	case r.Method == http.MethodGet && getPostRe.MatchString(r.URL.Path):
		{
			h.getPost(w,r)
			return
		}
	case r.Method == http.MethodGet && getUserPostRe.MatchString(r.URL.Path):
		{
			h.getUserPost(w,r)
			return
		}
	default:
		{
			notFound(w, r)
			return
		}
	}


}

//Create
		func (h *userHandler) createUser(w http.ResponseWriter, r *http.Request){
			var u models.User
			if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
				fmt.Fprintf(w, "error")
				return
			}

			key :=  os.Getenv("ENCRYPT_KEY")
			u.Password = encrypt.EncryptString([]byte(key), u.Password)

			insertResult, err := h.col.InsertOne(h.ctx, u)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				panic(err)
			}
			
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(insertResult)

		}

		func (h *postHandler) createPost(w http.ResponseWriter, r *http.Request){
			var p models.Post
			if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
				fmt.Fprintf(w, "error")
				return
			}

			insertResult, err := h.col.InsertOne(h.ctx, p)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				panic(err)
			}
			
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(insertResult)
		}

//Retrieve
		func (h *userHandler) getUser(w http.ResponseWriter, r *http.Request){

			userId := path.Base(r.URL.Path)
			id, _ := primitive.ObjectIDFromHex(userId)
			filter := bson.M{"_id": id}
			
			var result struct {}
			err := h.col.FindOne(h.ctx, filter).Decode(&result)
			if err != nil {
				log.Fatal(err)
			}

			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(result)
		}

		func (h *postHandler) getPost(w http.ResponseWriter, r *http.Request){

			postId := path.Base(r.URL.Path)
			id, _ := primitive.ObjectIDFromHex(postId)
			filter := bson.M{"_id": id}
			
			var result struct {}
			err := h.col.FindOne(h.ctx, filter).Decode(&result)
			if err != nil {
				log.Fatal(err)
			}
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(result)
		}

		func (h *postHandler) getUserPost(w http.ResponseWriter, r *http.Request){

			
			userId := path.Base(r.URL.Path)
			id, _ := primitive.ObjectIDFromHex(userId)
			filter := bson.M{"userId": id}
			
			var result struct {}
			cur, err := h.col.Find(h.ctx, filter)
			if err != nil {
				log.Fatal(err)
			}
			defer cur.Close(h.ctx)

			var results []bson.M
			if err := cur.All(context.TODO(), &results); err != nil {
			log.Fatal(err)
			}

			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(result)

		}

		func notFound(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("not found"))
		}


func RequestHandler(db *mongo.Database, ctx context.Context){
	userH := &userHandler{
		col: *db.Collection("users"),
		ctx: ctx,
	}
	postH := &postHandler{
		col: *db.Collection("posts"),
		ctx: ctx,
	}

	mux := http.NewServeMux()

	mux.Handle("/users/", userH)
	mux.Handle("/posts/", postH)
	mux.Handle("/posts/users/", postH)

	port := os.Getenv("PORT")
	log.Fatal(http.ListenAndServe(port, mux))
}

