package core

import . "gopkg.in/check.v1"

type SuiteBareJob struct{}

var _ = Suite(&SuiteBareJob{})

func (s *SuiteBareJob) TestGetters(c *C) {
	job := &BareJob{
		Name:     "foo",
		Schedule: "bar",
		Command:  "qux",
	}

	c.Assert(job.GetName(), Equals, "foo")
	c.Assert(job.GetSchedule(), Equals, "bar")
	c.Assert(job.GetCommand(), Equals, "qux")
}

func (s *SuiteBareJob) TestHistory(c *C) {
	eA := NewExecution()
	eB := NewExecution()

	job := &BareJob{}
	job.AddHistory(eA, eB)

	h := job.History()
	c.Assert(h, HasLen, 2)
	c.Assert(h[0], DeepEquals, eA)
	c.Assert(h[1], DeepEquals, eB)
}

func (s *SuiteBareJob) TestNotifyStartStop(c *C) {
	job := &BareJob{}

	job.NotifyStart()
	c.Assert(job.Running(), Equals, int32(1))

	job.NotifyStop()
	c.Assert(job.Running(), Equals, int32(0))
}
