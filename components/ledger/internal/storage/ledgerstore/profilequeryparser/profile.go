package profilequeryparser

type profilePattern map[string]ICriterion
type Profile map[string][]ICriterion 
type Stack []Profile

func (pp *profilePattern) MatchPatterns(key string) []ICriterion {
	temp := make([]ICriterion, 0)
	for _,criterion := range *pp{
		if(criterion.Check(key)){ temp = append(temp, criterion)}
	}
	return temp 
}

func (p *Profile) Add(c ICriterion){
	(*p)[c.Name()] = append((*p)[c.Name()], c)
}



func (s *Stack) Copy() *Stack{

	s2 := make(Stack, 0)
	
	for _,profile := range *s {

		new_profile := make(Profile)
		
		for name ,criterions := range profile {
			c2s := make([]ICriterion, 0)
			for _, criterion := range criterions {
				c2s = append(c2s, criterion.Copy())
				
			}
			new_profile[name] = c2s
		}		

		s2 = append(s2, new_profile)
	}
	return &s2

}

func (s *Stack) Add(c ICriterion){
	for _, profile := range *s {
		profile.Add(c.Copy())
	}
}


func NewProfilePattern(criterions ...*ICriterion) *profilePattern {
	pp := make(profilePattern)
	for _,criterion := range criterions{
		pp[(*criterion).Name()] = *criterion
	}
	return &pp 
}







