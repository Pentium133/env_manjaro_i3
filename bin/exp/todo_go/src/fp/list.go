package fp

type List struct {
  head Any
  tail *List
  functor Functor
}

var Nil List = List { nil, nil, EmptyFunctor }

func (l List) IsEmpty() bool {
  return l.head == nil && l.tail == nil
}

func (l List) IsNotEmpty() bool {
  return !l.IsEmpty()
}

/**
 * Copy (references only) list with empty functors
 */
func (l List) Copy() List {
  if l.IsEmpty() {
    return l
  } else {
    tail := l.tail.Copy()
    return List {
      head : l.head,
      tail : &tail,
      functor : EmptyFunctor,
    }
  }
}

/**
 * Add head to list
 *
 * O(1)
 */
func (l List) Cons(e Any) List {
  tail := l.Copy()
  xs := List {
    head: e,
    tail: &tail,
    functor: EmptyFunctor,
  }

  return xs
}

/**
 * Add lazy filter
 */
func (l List) Filter(p Predicate) List {
  return List {
    head : l.head,
    tail : l.tail,
    functor : func(e Any) Any {
      if processed := l.functor(e); processed != nil && p(processed) {
          return processed
      } else {
          return nil                                    //TODO return None
      }
    },
  }
}

/**
 * Add lazy functor
 */
func (l List) Map(f Functor) List {
  return List {
    head : l.head,
    tail : l.tail,
    functor : func(e Any) Any {
      if processed := l.functor(e); processed != nil {
        return f(processed)
      } else {
        return nil
      }
    },
  }
}

func (l List) Reverse() List {
  xs := Nil
  l.Foreach(func(e Any) {
    xs = xs.Cons(e)
  })
  return xs
}

func (l1 List) Zip(l2 List) List {
  zipped := Nil
  it1 := &l1
  it2 := &l2
  for {
    if (it1.head != nil && it2.head != nil) {
      zipped = zipped.Cons(Tuple2 { it1.head, it2.head })
      it1 = it1.tail
      it2 = it2.tail
    } else {
      break
    }
  }

  return zipped.Reverse()
}

/**
 * Materialize list: apply functors, filters before each element will be passed to the lambda
 */
func (l List) Foreach(f func(Any)) {
  if l.head != nil {
    processed := l.functor(l.head)
    if processed != nil {
      f(processed)
    }
  }

  if l.tail != nil {
    l.tail.mapHead(l.functor).Foreach(f)
  }
}


//-- internal --------------------------------------------------------------------

/**
 * Add lazy functor to the head only
 */
func (l List) mapHead(f Functor) List {
  return List {
    head: l.head,
    tail: l.tail,
    functor : func(e Any) Any {
      processed := l.functor(e)
      if processed == nil {
        return nil
      } else {
        return f(processed)
      }
    },
  }
}

