package nodeData

import (
    "github.com/pkg/errors"
)

type TxPool struct {
    queue []ReviewData
}

func NewTxPool() TxPool {
    newQueue := make([]ReviewData, 0)
    return TxPool{queue: newQueue}
}

func (q *TxPool) Offer(item ReviewData) {
    q.queue = append(q.queue, item)
}

func (q *TxPool) Poll() (ReviewData, error) {
    if len(q.queue) == 0 {
        return ReviewData{}, errors.New("Queue is empty ")
    }
    res := q.queue[0]
    q.queue = q.queue[1:]
    return res, nil
}

/**
TxPool is Empty
 */
func (q *TxPool) IsEmpty() bool {
    if len(q.queue) == 0 {
        return true
    }
    return false
}