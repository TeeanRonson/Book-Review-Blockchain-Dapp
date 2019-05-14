package Utilities

import (
    "github.com/pkg/errors"
    "github.com/teeanronson/cs686-blockchain-p3-TeeanRonson/p5/nodeData"
)

type Queue struct {
    queue []nodeData.ReviewData
}

func NewQueue() Queue {
    newQueue := make([]nodeData.ReviewData, 0)
    return Queue{queue: newQueue}
}

func (q *Queue) Offer(item nodeData.ReviewData) {
    q.queue = append(q.queue, item)
}

func (q *Queue) Poll() (nodeData.ReviewData, error) {
    if len(q.queue) == 0 {
        return nodeData.ReviewData{}, errors.New("poll from an empty queue")
    }
    res := q.queue[0]
    q.queue = q.queue[1:]
    return res, nil
}