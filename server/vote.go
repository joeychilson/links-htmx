package server

import (
	"fmt"
	"net/http"

	"github.com/go-chi/httplog/v2"
	"github.com/google/uuid"

	"github.com/joeychilson/links/components/comment"
	"github.com/joeychilson/links/components/link"
	"github.com/joeychilson/links/database"
)

func (s *Server) Vote() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		oplog := httplog.LogEntry(ctx)
		user := s.UserFromContext(ctx)

		linkID := r.URL.Query().Get("link_id")
		commentID := r.URL.Query().Get("comment_id")
		voteDir := r.URL.Query().Get("vote")
		redirect := r.URL.Query().Get("redirect")

		if linkID == "" {
			w.Header().Set("HX-Redirect", redirect)
			w.WriteHeader(http.StatusOK)
			return
		}

		linkUUID, err := uuid.Parse(linkID)
		if err != nil {
			oplog.Error("failed to parse link id", "error", err)
			w.Header().Set("HX-Redirect", redirect)
			w.WriteHeader(http.StatusOK)
			return
		}

		var vote int32
		if voteDir == "up" {
			vote = 1
		} else if voteDir == "down" {
			vote = -1
		} else {
			w.Header().Set("HX-Redirect", redirect)
			w.WriteHeader(http.StatusOK)
			return
		}

		if commentID != "" {
			commentUUID, err := uuid.Parse(commentID)
			if err != nil {
				oplog.Error("failed to parse comment id", "error", err)
				w.Header().Set("HX-Redirect", redirect)
				w.WriteHeader(http.StatusOK)
				return
			}

			err = s.queries.CommentVote(ctx, database.CommentVoteParams{
				UserID:    user.ID,
				CommentID: commentUUID,
				Vote:      vote,
			})
			if err != nil {
				oplog.Error("failed to vote on comment", "error", err)
				w.Header().Set("HX-Redirect", fmt.Sprintf("/link?id=%s", linkID))
				w.WriteHeader(http.StatusOK)
				return
			}

			commentRow, err := s.queries.Comment(ctx, database.CommentParams{
				UserID:    user.ID,
				CommentID: commentUUID,
			})
			if err != nil {
				oplog.Error("failed to get comment", "error", err)
				w.Header().Set("HX-Redirect", fmt.Sprintf("/link?id=%s", linkID))
				w.WriteHeader(http.StatusOK)
				return
			}

			oplog.Info("user voted on comment", "comment_id", commentUUID)
			comment.Component(comment.Props{User: user, Comment: commentRow}).Render(ctx, w)
			return
		} else {
			err = s.queries.LinkVote(ctx, database.LinkVoteParams{
				UserID: user.ID,
				LinkID: linkUUID,
				Vote:   vote,
			})
			if err != nil {
				oplog.Error("failed to vote on link", "error", err)
				w.Header().Set("HX-Redirect", redirect)
				w.WriteHeader(http.StatusOK)
				return
			}

			linkRow, err := s.queries.Link(ctx, database.LinkParams{
				UserID: user.ID,
				LinkID: linkUUID,
			})
			if err != nil {
				oplog.Error("failed to get link", "error", err)
				w.Header().Set("HX-Redirect", redirect)
				w.WriteHeader(http.StatusOK)
				return
			}

			oplog.Info("user voted on link", "link_id", linkID)
			link.Component(link.Props{User: user, Link: linkRow}).Render(ctx, w)
			return
		}
	}
}